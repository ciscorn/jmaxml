import pathlib

from jmx_codegen import load_schema
from jmx_codegen.langs.go import GoGenerator
from jmx_codegen.langs.typescript import TypeScriptGenerator


if __name__ == "__main__":
    # XML Schema を読み込む
    src_path = pathlib.Path(__file__).parent.parent.joinpath("./assets/schema")
    schema = load_schema(str(src_path))

    # Golang 用の型定義を生成する
    dst_path = pathlib.Path(__file__).parent.parent.joinpath("./jmaxml-go")
    GoGenerator(schema).generate(str(dst_path))

    # TypeScript 用の型定義を生成する
    dst_path = pathlib.Path(__file__).parent.parent.joinpath("./jmaxml-ts")
    TypeScriptGenerator(schema).generate(str(dst_path))
