"""Go用の型定義を出力します"""

import io
import os.path

from collections import Counter
from typing import cast

from jmx_codegen.langs.common import pluralize
from jmx_codegen.types import (
    XsAttribute,
    XsElement,
    XsBase,
    XsEnumeration,
    XsPrimitive,
    XsSchema,
    XsComplexType,
    XsChildElement,
)


_primitive_map = {
    "StringList": "StringList",
    "jmx_eb:nullablefloat": "NullableFloat",
    "jmx_eb:nullableinteger": "NullableInteger",
    "xs:anyURI": "string",
    "xs:boolean": "bool",
    "xs:dateTime": "time.Time",
    "xs:dateTime-nillable": "NullableDateTime",
    "xs:duration": "Duration",
    "xs:float": "float64",
    "xs:gMonthDay": "string",
    "xs:int": "int32",
    "xs:string": "string",
    "xs:token": "string",
    "xs:unsignedByte": "uint8",
}

_prefix_map = {
    "jmx": "http://xml.kishou.go.jp/jmaxml1/",
    "jmx_add": "http://xml.kishou.go.jp/jmaxml1/addition1/",
    "jmx_eb": "http://xml.kishou.go.jp/jmaxml1/elementBasis1/",
    "jmx_seis": "http://xml.kishou.go.jp/jmaxml1/body/seismology1/",
    "jmx_volc": "http://xml.kishou.go.jp/jmaxml1/body/volcanology1/",
    "jmx_mete": "http://xml.kishou.go.jp/jmaxml1/body/meteorology1/",
    "jmx_ib": "http://xml.kishou.go.jp/jmaxml1/informationBasis1/",
}


class GoGenerator:
    def __init__(self, schema: XsSchema):
        self.schema = schema

    def generate(self, dir: str):
        schema = self.schema

        # 型定義をもつ名前空間を収集する
        prefixes = Counter()
        for (name, item) in schema.type_map.items():
            if isinstance(item, XsComplexType):
                prefix = name.split(":")[0]
                prefixes[prefix] += 1

        for prefix in prefixes:
            # 名前空間ごとに types_{prefix}.go ファイルを生成する
            with open(os.path.join(dir, f"types_{prefix}.go"), "w", encoding="utf-8") as f:
                f.write("""// Code generated by jmx_codegen; DO NOT EDIT.\n\n""")
                f.write("""package jmaxml\n\nimport "time"\n\n""")
                for (name, item) in schema.type_map.items():
                    if name.startswith(prefix + ":"):
                        if isinstance(item, XsComplexType):
                            self._write_type(f, "", item)
                            f.write("\n\n")

    def _get_modifier(self, child: XsChildElement) -> str:
        if child.max_occurs is None:
            if child.type in _primitive_map:
                return "[]"
            else:
                return "[]*"
        elif child.max_occurs == 1:
            if child.min_occurs == 0:
                return "*"
            elif child.min_occurs == 1:
                return ""
            else:
                raise RuntimeError("min_occurs must be 0 or 1")
        else:
            raise RuntimeError("max_occurs must be None or 1")

    def _get_attrib_modifier(self, attr: XsAttribute) -> str:
        if attr.use == "optional":
            return "*"
        elif attr.use == "required":
            return ""
        else:
            raise RuntimeError("attribute's use must be required or optional")

    def _to_field_name(self, name: str, plural: bool) -> str:
        """Schema上の要素名をGoのフィールド名に変換する"""

        if plural:
            name = pluralize(name)

        if ":" in name:
            prefix, name = name.split(":", 1)
        else:
            prefix = ""
        name = name[:1].upper() + name[1:]
        if name == "Body":
            # Body要素についてのみ、jmx_mete:Body, jmx_seis:Body, jmx_volc:Body の名前が衝突するため、
            # フィールド名を MeteBody, SeisBody, VolcBody とする。 
            return prefix.split("_")[-1].capitalize() + name
        else:
            return name

    def _to_type_name(self, name: str) -> str:
        """Schema上の型名をGoの型名に変換する"""

        prefix, name = name.split(":", 1)
        # 'jmx_mete' などから 'jmx_' を除去する
        parts = prefix.split("_", 1)
        assert parts[0] == "jmx"
        parts = parts[1:]
        # 'type.' を除去して先頭文字を大文字にする
        if name.startswith("type."):
            name = name[5:]
            name = name[:1].upper() + name[1:]
        return "".join(p.capitalize() for p in parts) + name

    def _to_xml_tag(self, name: str, omittable: bool, plural: bool) -> str:
        """Schema上の要素をGoのXMLタグとJSONタグ名に変換する"""

        prefix, name = name.split(":", 1)
        jsonname = name[:1].lower() + name[1:]
        if plural:
            jsonname = pluralize(jsonname)
        if name == "Body":
            # Body要素のみ、XML名前空間を考慮する
            uri = _prefix_map[prefix]
            return f'`xml:"{uri} {name}" json:"{prefix}{name},omitempty"`'
        elif omittable:
            return f'`xml:"{name}" json:"{jsonname},omitempty"`'
        else:
            return f'`xml:"{name}" json:"{jsonname}"`'

    def _write_type(self, f: io.TextIOBase, indent: str, _type: XsBase):
        schema = self.schema

        if isinstance(_type, XsElement):
            pass
        elif isinstance(_type, XsComplexType):
            if len(indent) > 0:
                f.write(self._to_type_name(_type.name))
                return

            f.write(f"type {self._to_type_name(_type.name)} ")
            if not (_type.attributes or _type.elements):
                assert _type.content_type is not None
                f.write(f"{_type.content_type}")
            else:
                indent = indent + "\t"
                f.write("struct {\n")
                if _type.content_type:
                    f.write(f"{indent}Value  ")
                    self._write_type(f, indent, schema.type_map[_type.content_type])
                    f.write('  `xml:",chardata"`\n')
                for name, attr in _type.attributes.items():
                    f.write(f"{indent}{self._to_field_name(name, plural=False)}  ")
                    f.write(self._get_attrib_modifier(attr))
                    self._write_type(f, indent + "  ", schema.type_map[attr.type])
                    name = name.split(":", 1)[-1]
                    if attr.use == "optional":
                        f.write(f'  `xml:"{name},attr" json:"{name},omitempty"`')
                    else:
                        f.write(f'  `xml:"{name},attr" json:"{name}"`')
                    f.write("\n")

                for child in _type.elements:
                    modifier = self._get_modifier(child)
                    omittable = child.min_occurs == 0
                    plural = child.max_occurs is None
                    if child.type:
                        assert child.name is not None
                        f.write(f"{indent}{self._to_field_name(child.name, plural=plural)}  {modifier}")
                        self._write_type(f, indent + "  ", schema.type_map[child.type])
                        f.write(self._to_xml_tag(child.name, omittable, plural=plural) + '\n')
                    elif child.ref:
                        ref = cast(XsElement, schema.type_map[child.ref])
                        field_name = self._to_field_name(ref.name, plural=plural)
                        type_name = self._to_type_name(ref.type)
                        f.write(f"{indent}{field_name}  {modifier}{type_name}")
                        f.write(self._to_xml_tag(ref.name, omittable, plural=plural) + '\n')
                    else:
                        raise RuntimeError("element must have type or ref")
                f.write(f"{indent[:-2]}}}")

        elif isinstance(_type, XsPrimitive):
            f.write(f"{_primitive_map[_type.type]}")
        elif isinstance(_type, XsEnumeration):
            f.write(f"{_primitive_map[_type.type]}")
        else:
            raise RuntimeError("unknown type")
