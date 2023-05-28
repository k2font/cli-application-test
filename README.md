# cli-application-test
## Features
- 与えた文字列を反転するためのCLIアプリケーション
- GoでCLIアプリケーションを開発する練習のために作成
- ライブラリにCobraを採用しました
  - コードが書きやすく、ドキュメントがわかりやすいため + デファクトスタンダードっぽく感じたため

## Requirements
- spf13/cobra@1.7.0

## Usage
- `go build && ./cli-applicaion-test ...`

```bash
🍎 ./cli-application-test -h
このCLIアプリケーションは、ShoichiroがCobraの使い方を学ぶために書いた習作です

Usage:
  stringer [flags]
  stringer [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  inspect     Inspects a string
  reverse     Reverse a string

Flags:
  -h, --help      help for stringer
  -v, --version   version for stringer

Use "stringer [command] --help" for more information about a command.
```