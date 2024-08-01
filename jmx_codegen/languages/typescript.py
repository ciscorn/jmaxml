"""TypeScript用の型定義を出力します (Experimental)"""

import io
import os.path
from typing import cast

from jmx_codegen.languages.common import pluralize
from jmx_codegen.types import (
    XsAttribute,
    XsBase,
    XsChildElement,
    XsComplexType,
    XsElement,
    XsEnumeration,
    XsPrimitive,
    XsSchema,
)

_PRIMITIVE_MAP = {
    "StringList": "string[]",
    "jmx_eb:nullablefloat": "number | null",
    "jmx_eb:nullableinteger": "number | null",
    "xs:anyURI": "string",
    "xs:boolean": "boolean",
    "xs:dateTime": "RFC3339String",
    "xs:dateTime-nillable": "RFC3339String | null",
    "xs:duration": "Duration",
    "xs:float": "number",
    "xs:gMonthDay": "string",
    "xs:int": "number",
    "xs:string": "string",
    "xs:token": "string",
    "xs:unsignedByte": "number",
}

_CUSTOM_TYPES = """
type RFC3339String = string

// ISO8601 の Duration を表す
export type Duration = {
  y?: number   // Year
  m?: number   // Month
  w?: number   // Week
  d?: number   // Day
  th?: number  // Hour
  tm?: number  // Minute
  ts?: number  // Second
}
"""


class TypeScriptGenerator:
    def __init__(self, schema: XsSchema):
        self.schema = schema

    def generate(self, dir: str):
        schema = self.schema

        with open(os.path.join(dir, "jmaxml.ts"), "w", encoding="utf-8") as f:
            f.write("""// Code generated by jmx_codegen; DO NOT EDIT.\n\n""")
            for _, item in sorted(schema.type_map.items()):
                if isinstance(item, XsComplexType):
                    self._write_type(f, "", item)
                    f.write("\n\n")

            f.write(_CUSTOM_TYPES)

    def _get_type_modifier(self, child: XsChildElement) -> str:
        if child.max_occurs is None:
            return "[]"
        elif child.max_occurs == 1:
            return ""
        else:
            raise RuntimeError("max_occurs must be None or 1")

    def _get_field_modifier(self, child: XsChildElement) -> str:
        if child.min_occurs == 0:
            if child.max_occurs is None:
                return ""
            elif child.max_occurs == 1:
                return "?"
            else:
                raise RuntimeError("min_occurs must be 0 or 1")
        elif child.min_occurs == 1:
            return ""
        else:
            raise RuntimeError("min_occurs must be 0 or 1")

    def _get_attrib_modifier(self, attr: XsAttribute) -> str:
        if attr.use == "optional":
            return "?"
        elif attr.use == "required":
            return ""
        else:
            raise RuntimeError("attribute's use must be required or optional")

    def _to_field_name(self, name: str, plural: bool) -> str:
        """Schema上の要素名をTypeScriptのフィールド名に変換する"""
        if plural:
            name = pluralize(name)

        if ":" in name:
            prefix, name = name.split(":", 1)
        else:
            prefix = ""
        if name == "Body":
            # Body要素についてのみ、
            # jmx_mete:Body, jmx_seis:Body, jmx_volc:Body の名前が衝突するため、
            # フィールド名を MeteBody, SeisBody, VolcBody とする。
            return prefix.split("_")[-1] + name
        else:
            return name[:1].lower() + name[1:]

    def _to_type_name(self, name: str) -> str:
        """Schema上の型名をTypeScriptの型名に変換する"""
        prefix, name = name.split(":", 1)
        parts = prefix.split("_", 1)
        # 'jmx' を除去する
        assert parts[0] == "jmx"
        parts = parts[1:]
        # 'type.' を除去して先頭文字を大文字にする
        if name.startswith("type."):
            name = name[5:]
            name = name[:1].upper() + name[1:]
        return "JMX" + "".join(p.capitalize() for p in parts) + name

    def _write_type(self, f: io.TextIOBase, indent: str, _type: XsBase):
        schema = self.schema

        if isinstance(_type, XsElement):
            pass
        elif isinstance(_type, XsComplexType):
            if len(indent) > 0:
                f.write(self._to_type_name(_type.name))
                return

            f.write(f"export type {self._to_type_name(_type.name)} = ")
            if not (_type.attributes or _type.elements):
                assert _type.content_type is not None
                f.write(f"{_type.content_type}")
            else:
                indent = indent + "  "
                f.write("{\n")
                if _type.content_type:
                    f.write(f"{indent}value: ")
                    self._write_type(f, indent, schema.type_map[_type.content_type])
                    f.write("\n")
                for name, attr in _type.attributes.items():
                    attr_modifier = self._get_attrib_modifier(attr)
                    f.write(
                        f"{indent}{self._to_field_name(name, plural=False)}{attr_modifier}: "
                    )
                    self._write_type(f, indent + "  ", schema.type_map[attr.type])
                    name = name.split(":", 1)[-1]
                    f.write("\n")

                for child in _type.elements:
                    type_modifier = self._get_type_modifier(child)
                    field_modifier = self._get_field_modifier(child)
                    plural = child.max_occurs is None
                    if child.type:
                        assert child.name is not None
                        f.write(
                            f"{indent}{self._to_field_name(child.name, plural=plural)}{field_modifier}: "
                        )
                        self._write_type(f, indent + "  ", schema.type_map[child.type])
                        f.write(type_modifier)
                    elif child.ref:
                        ref = cast(XsElement, schema.type_map[child.ref])
                        field_name = self._to_field_name(ref.name, plural=plural)
                        type_name = self._to_type_name(ref.type)
                        f.write(
                            f"{indent}{field_name}{field_modifier}: {type_name}{type_modifier}"
                        )
                    else:
                        raise RuntimeError("element must have type or ref")
                    f.write("\n")
                f.write(f"{indent[:-2]}}}")

        elif isinstance(_type, XsPrimitive):
            f.write(f"{_PRIMITIVE_MAP[_type.type]}")
        elif isinstance(_type, XsEnumeration):
            values = " | ".join(f'"{v}"' for v in _type.values)
            f.write(values)
        else:
            raise RuntimeError("unknown type")