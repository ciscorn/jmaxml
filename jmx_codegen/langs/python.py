"""Python 用の型定義を出力します"""

import re
import io
import os.path
from graphlib import TopologicalSorter

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


_PRIMITIVE_TYPES = {
    "StringList": "list[str]",
    "jmx_eb:nullablefloat": "Optional[float]",
    "jmx_eb:nullableinteger": "Optional[int]",
    "xs:anyURI": "str",
    "xs:boolean": "bool",
    "xs:dateTime": "datetime",
    "xs:dateTime-nillable": "Optional[datetime]",
    "xs:duration": "Duration",
    "xs:float": "float",
    "xs:gMonthDay": "str",
    "xs:int": "int",
    "xs:string": "str",
    "xs:token": "str",
    "xs:unsignedByte": "int",
}

_PRIMITIVE_COERCERS = {
    "StringList": "parse_stringlist",
    "jmx_eb:nullablefloat": "float",
    "jmx_eb:nullableinteger": "int",
    "xs:dateTime": "parse_datetime",
    "xs:dateTime-nillable": "parse_datetime",
    "xs:duration": "from_iso8601_duration",
}

_RESERVED_WORDS = {
    "from",
}

_HEADER = """# Code generated by jmx_codegen; DO NOT EDIT.

from __future__ import annotations
from typing import Optional
from datetime import datetime

from .utils import ElementBase, child, childtext, attribute, text
from .custom_types import parse_stringlist, parse_datetime, Duration

from_iso8601_duration = Duration.from_iso8601_duration
"""


def camel_to_snake(name):
    name = re.sub("(.)([A-Z][a-z]+)", r"\1_\2", name)
    return re.sub("([a-z0-9])([A-Z])", r"\1_\2", name).lower()


class PythonGenerator:
    def __init__(self, schema: XsSchema):
        self.schema = schema

    def generate(self, dir: str):
        schema = self.schema

        # 型名をトポロジカルソートする
        graph = {}
        for (name, item) in schema.type_map.items():
            if isinstance(item, XsComplexType):
                depends = [depend.type for depend in item.elements if depend.type]
                depends += [
                    depend.ref.replace(":", ":type.")
                    for depend in item.elements
                    if depend.ref
                ]
                graph[name] = depends

        types_topo = TopologicalSorter(graph).static_order()

        with open(os.path.join(dir, "jmx_types.py"), "w", encoding="utf-8") as f:
            f.write(_HEADER)
            f.write("\n\n")
            for name in types_topo:
                if name is None:
                    continue

                if item := schema.type_map.get(name):
                    if isinstance(item, XsComplexType):
                        f.write(self._write_type("", item, is_typehint=True))
                        f.write("\n\n")

    def _get_modifier(self, child: XsChildElement, type_repr: str) -> str:
        if child.max_occurs is None:
            return f"list[{type_repr}]"
        elif child.max_occurs == 1:
            if child.min_occurs == 0:
                return f"Optional[{type_repr}]"
            elif child.min_occurs == 1:
                return type_repr
            else:
                raise RuntimeError("min_occurs must be 0 or 1")
        else:
            raise RuntimeError("max_occurs must be None or 1")

    def _get_attrib_modifier(self, attr: XsAttribute, type_repr: str) -> str:
        if attr.use == "optional":
            return f"Optional[{type_repr}]"
        elif attr.use == "required":
            return type_repr
        else:
            raise RuntimeError("attribute's use must be required or optional")

    def _to_field_name(self, name: str, plural: bool) -> str:
        """Schema上の要素名をPythonのフィールド名に変換する"""
        original_name = name

        if plural:
            name = pluralize(name)

        if ":" in name:
            prefix, name = name.split(":", 1)
        else:
            prefix = ""

        name = camel_to_snake(name[:1].upper() + name[1:])
        if original_name in ["jmx_mete:Body", "jmx_seis:Body", "jmx_volc:Body"]:
            # Body要素についてのみ、jmx_mete:Body, jmx_seis:Body, jmx_volc:Body の名前が衝突するため、
            # フィールド名を mete_body, seis_body, volc_body とする。
            return prefix.split("_")[-1] + "_" + name
        elif name in _RESERVED_WORDS:
            return name + "_"
        else:
            return name

    def _to_type_name(self, name: str) -> str:
        """Schema上の型名をPythonの型名に変換する"""

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

    def _to_xml_tag(self, name: str, omittable: bool, _ty, plural: bool) -> str:
        """Schema上の要素をPythonのXMLタグとJSONタグ名に変換する"""

        primitive = isinstance(_ty, (XsPrimitive, XsEnumeration))
        _type_name = self._write_type("  ", _ty, is_typehint=False)
        many = ", many=True" if plural else ""

        if not _type_name:
            raise RuntimeError(f"unknown type: {_ty}")
        elif omittable:
            if primitive:
                return f' = childtext({_type_name}, "{name}")'
            else:
                return f' = child({_type_name}, "{name}"{many})'
        else:
            if primitive:
                return f' = childtext({_type_name}, "{name}")'
            else:
                return f' = child({_type_name}, "{name}"{many})'

    def _write_type(self, indent: str, _type: XsBase, is_typehint: bool) -> str:
        f = io.StringIO()
        schema = self.schema

        if isinstance(_type, XsElement):
            return self._to_type_name(_type.type)
        elif isinstance(_type, XsComplexType):
            if len(indent) > 0:
                f.write(self._to_type_name(_type.name))
                return f.getvalue()

            f.write(f"class {self._to_type_name(_type.name)}(ElementBase):\n")
            if not (_type.attributes or _type.elements):
                assert _type.content_type is not None
                f.write(f"{_type.content_type}")
            else:
                indent = indent + "    "
                if _type.content_type:
                    _type_hint = self._write_type(
                        indent, schema.type_map[_type.content_type], is_typehint=True
                    )
                    _ty = self._write_type(
                        indent, schema.type_map[_type.content_type], is_typehint=False
                    )
                    f.write(f"{indent}content: {_type_hint} = text({_ty})\n")
                for name, attr in _type.attributes.items():
                    f.write(f"{indent}{self._to_field_name(name, plural=False)}: ")
                    _t = self._write_type(
                        indent + "  ", schema.type_map[attr.type], is_typehint=True
                    )
                    f.write(self._get_attrib_modifier(attr, _t))
                    name = name.split(":", 1)[-1]
                    if attr.use == "optional":
                        f.write(f' = attribute({_t}, "{name}")')
                    else:
                        f.write(f' = attribute({_t}, "{name}")')
                    f.write("\n")

                for child in _type.elements:
                    omittable = child.min_occurs == 0
                    plural = child.max_occurs is None
                    if child.type:
                        assert child.name is not None
                        _ty = schema.type_map[child.type]
                        _t = self._get_modifier(
                            child,
                            self._write_type(indent + "  ", _ty, is_typehint=True),
                        )
                        f.write(
                            f"{indent}{self._to_field_name(child.name, plural=plural)}: {_t}"
                        )
                        f.write(
                            self._to_xml_tag(child.name, omittable, _ty, plural=plural)
                            + "\n"
                        )
                    elif child.ref:
                        ref = cast(XsElement, schema.type_map[child.ref])
                        field_name = self._to_field_name(ref.name, plural=plural)
                        type_name = self._to_type_name(ref.type)
                        _t = self._get_modifier(child, type_name)
                        f.write(f"{indent}{field_name}: {_t}")
                        f.write(
                            self._to_xml_tag(ref.name, omittable, ref, plural=plural)
                            + "\n"
                        )
                    else:
                        raise RuntimeError("element must have type or ref")

        elif isinstance(_type, (XsPrimitive, XsEnumeration)):
            if is_typehint:
                f.write(f"{_PRIMITIVE_TYPES[_type.type]}")
            else:
                if _type.type in _PRIMITIVE_COERCERS:
                    f.write(f"{_PRIMITIVE_COERCERS[_type.type]}")
                else:
                    f.write(f"{_PRIMITIVE_TYPES[_type.type]}")
        else:
            raise RuntimeError("unknown type")

        return f.getvalue()