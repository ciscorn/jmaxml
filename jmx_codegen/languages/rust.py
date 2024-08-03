"""Rust用の型定義を出力します"""

import io
import os.path
from typing import cast

from jmx_codegen.languages.common import pluralize
from jmx_codegen.types import (
    XsBase,
    XsChildElement,
    XsComplexType,
    XsElement,
    XsEnumeration,
    XsPrimitive,
    XsSchema,
)

from .common import camel_to_snake

_HEADER = """
use serde::{Deserialize, Serialize};

pub use super::builtins::*;
"""

_PRIMITIVE_MAP = {
    "StringList": "StringList",
    "jmx_eb:nullablefloat": "Option<f64>",
    "jmx_eb:nullableinteger": "Option<i64>",
    "xs:anyURI": "String",
    "xs:boolean": "bool",
    "xs:dateTime": "DateTime",
    "xs:dateTime-nillable": "NullableDateTime",
    "xs:duration": "Duration",
    "xs:float": "f64",
    "xs:gMonthDay": "String",
    "xs:int": "i64",
    "xs:string": "String",
    "xs:token": "String",
    "xs:unsignedByte": "u8",
}


def _json_name(name: str, plural=False) -> str:
    if plural:
        name = pluralize(name)
    return name[:1].lower() + name[1:]


class RustGenerator:
    def __init__(self, schema: XsSchema):
        self.schema = schema

    def generate(self, dir: str):
        schema = self.schema

        with open(
            os.path.join(dir, "src", "model", "generated.rs"), "w", encoding="utf-8"
        ) as f:
            f.write("""// Code generated by jmx_codegen; DO NOT EDIT.\n\n""")
            f.write(_HEADER)
            f.write("\n")

            for name, item in sorted(schema.type_map.items(), key=lambda x: x[0]):
                if isinstance(item, XsComplexType):
                    if name == "jmx:type.report":
                        # Report 型は自前で用意する (Bodyの扱いが特殊なため)
                        continue

                    self._write_type(f, "", item)
                    f.write("\n\n")

    def _get_modifier(self, child: XsChildElement) -> tuple[str, str]:
        if child.max_occurs is None:
            return ("Vec<", ">")
        elif child.max_occurs == 1:
            if child.min_occurs == 0:
                return ("Option<", ">")
            elif child.min_occurs == 1:
                return ("", "")
            else:
                raise RuntimeError("min_occurs must be 0 or 1")
        else:
            raise RuntimeError("max_occurs must be None or 1")

    def _to_field_name(self, name: str, plural: bool) -> str:
        """Schema上の要素名をRustのフィールド名に変換する"""

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
            name = prefix.split("_")[-1].capitalize() + name

        name = camel_to_snake(name)
        if name == "type":
            return "ty"
        return name

    def _to_type_name(self, name: str) -> str:
        """Schema上の型名をRustの型名に変換する"""

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

    def _write_type(self, f: io.TextIOBase, indent: str, _type: XsBase):
        schema = self.schema

        if isinstance(_type, XsElement):
            pass
        elif isinstance(_type, XsComplexType):
            if len(indent) > 0:
                f.write(self._to_type_name(_type.name))
                return

            f.write("#[derive(Debug, Serialize, Deserialize)]\n")
            f.write(f"pub struct {self._to_type_name(_type.name)} ")
            if not (_type.attributes or _type.elements):
                assert _type.content_type is not None
                f.write(f"{_type.content_type}")
            else:
                indent = indent + "    "
                f.write("{\n")
                if _type.content_type:
                    f.write(
                        '#[serde(rename(deserialize="$text", serialize="value"))]\n'
                    )
                    f.write(f"{indent}value: ")
                    if _type.content_type == "xs:string":
                        f.write("Option<String>")
                    else:
                        self._write_type(f, indent, schema.type_map[_type.content_type])
                    f.write(",\n")

                # attributes
                for name, attr in _type.attributes.items():
                    optional = attr.use == "optional"
                    elem_name = name.split(":", 1)[-1]
                    if optional:
                        f.write(
                            f'{indent}#[serde(rename(deserialize="@{elem_name}", serialize="@{_json_name(elem_name)}"), skip_serializing_if="Option::is_none")]\n'
                        )
                    else:
                        f.write(
                            f'{indent}#[serde(rename(deserialize="@{elem_name}", serialize="@{_json_name(elem_name)}"))]\n'
                        )

                    f.write(f"{indent}pub {self._to_field_name(name, plural=False)}: ")
                    if optional:
                        f.write("Option<")

                    if schema.type_map[attr.type].name == "StringList":  # Note: hack
                        f.write("Vec<String>")
                    else:
                        self._write_type(f, indent + "  ", schema.type_map[attr.type])

                    if optional:
                        f.write(">")
                    f.write(",\n")

                for child in _type.elements:
                    m_prefix, m_suffix = ("", "")
                    if child.max_occurs is None:
                        (m_prefix, m_suffix) = ("Vec<", ">")
                    elif child.max_occurs == 1:
                        if child.min_occurs == 0:
                            (m_prefix, m_suffix) = ("Option<", ">")
                        elif child.min_occurs == 1:
                            (m_prefix, m_suffix) = ("", "")
                        else:
                            raise RuntimeError("min_occurs must be 0 or 1")
                    else:
                        raise RuntimeError("max_occurs must be None or 1")

                    m_prefix, m_suffix = self._get_modifier(child)
                    plural = child.max_occurs is None
                    serde_attrs = ""

                    if m_prefix == "Option<":
                        serde_attrs += ', skip_serializing_if="Option::is_none"'
                    elif m_prefix == "Vec<":
                        serde_attrs += ', skip_serializing_if="Vec::is_empty", default'

                    if child.type:
                        assert child.name is not None
                        elem_name = child.name.split(":", 1)[-1]
                        f.write(
                            f'{indent}#[serde(rename(deserialize="{elem_name}", serialize="{_json_name(elem_name, plural)}"){serde_attrs})]\n'
                        )
                        f.write(
                            f"{indent}pub {self._to_field_name(child.name, plural=plural)}: "
                        )
                        f.write(m_prefix)
                        self._write_type(
                            f, indent + "    ", schema.type_map[child.type]
                        )
                        f.write(m_suffix)
                        f.write(",\n")
                    elif child.ref:
                        ref = cast(XsElement, schema.type_map[child.ref])
                        elem_name = ref.name.split(":", 1)[-1]
                        f.write(
                            f'{indent}#[serde(rename(deserialize="{elem_name}", serialize="{_json_name(elem_name, plural)}"){serde_attrs})]\n'
                        )
                        field_name = self._to_field_name(ref.name, plural=plural)
                        type_name = self._to_type_name(ref.type)
                        f.write(
                            f"{indent}pub {field_name}: {m_prefix}{type_name}{m_suffix}"
                        )
                        f.write(",\n")
                    else:
                        raise RuntimeError("element must have type or ref")
                f.write(f"{indent[:-2]}}}")

        elif isinstance(_type, XsPrimitive):
            f.write(f"{_PRIMITIVE_MAP[_type.type]}")
        elif isinstance(_type, XsEnumeration):
            f.write(f"{_PRIMITIVE_MAP[_type.type]}")
        else:
            raise RuntimeError("unknown type")