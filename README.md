# uq - Universal (de)serializer

[![Go](https://github.com/solarkennedy/uq/actions/workflows/go.yml/badge.svg)](https://github.com/solarkennedy/uq/actions/workflows/go.yml)

```
Usage:
  uq [-v] [-s FORMAT] [-t FORMAT] [FILE]

Options:
  -s <FORMAT>, --source <FORMAT>  Specify input format. [default: auto]
  -t <FORMAT>, --target <FORMAT>  Specify output format. [default: json]
  -v, --verbose                   Be more verbose [default: false]
  -h, --help     Show this screen
  --version     Show version

Formats:
  * json
  * yaml|yml
  * ini
  * toml
  * xml (Note: xml won't be a perfect conversion)
