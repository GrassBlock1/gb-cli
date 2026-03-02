# gb
## Motivation
I wanted to decode base* strings directly without using `echo`, so I created this tool with the basic operation for base* strings.

Soon I find that I can get some more knowledge upon Golang by coding on the tool, so I started to create more subcommands for my usage. And I tried to create this tool without any what is called "vibe coding", only with the help of search and chatbots. I make sure I understand most of the code when I do a copy-paste.
## Installation
```bash
git clone https://git.gay/gb/gb
go install .
# or
go build .
```
## Usage
```
➜ gb
gb is a cli tool for everyday tasks - base32/64/85 encoding/decoding, getting quotes, etc

Usage:
  gb [flags]
  gb [command]

Available Commands:
  b32         Encode/decode base32 string
  b64         Encode/decode base64 string
  b85         Encode/decode standard base85(ascii85) string
  bf          run a brainfuck script
  completion  Generate the autocompletion script for the specified shell
  get         Get things from the Internet
  help        Help about any command
  matrix      Get information of a matrix homeserver
  passwd      Generate a memorable secure password
  rand        Get real random number from multiple sources

Flags:
  -h, --help   help for gb

Use "gb [command] --help" for more information about a command.
```
## Notes
`gb` is originally named as an abbreviation for my nickname on the Internet, but it can be anything, at least I think so.

You may also find this tool at [codeberg](https://codeberg.org/grassblock/gb) and [GitHub](https://github.com/GrassBlock1/gb).
## License
This tool is under [Mozilla Public License, Version 2.0](https://www.mozilla.org/en-US/MPL/), see [LICENSE file](./LICENSE) for details.
