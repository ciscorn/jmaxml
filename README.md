# jmaxml

TODO: CI


(WIP) 気象庁防災情報XMLを型付きで読むためのライブラリです。今のところGoとTypeScript用の型定義のみです。

気象庁が提供する XML Schema をもとにコードを生成しています。

## Usage

各言語用ライブラリのREADMEを参照してください。

- [Go](./jmaxml-go/)
- [TypeScript](./jmaxml-ts/)

## Development

コードジェネレータは `./jmx_codegen/` ディレクトリ内で実装されています。

```bash
# コードジェネレータの再実行
make update

# テストの実行
make test
```
