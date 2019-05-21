# sync-dotenv

[![Build Status](https://wdp9fww0r9.execute-api.us-west-2.amazonaws.com/production/badge/atrox/sync-dotenv?style=flat-square)](https://wdp9fww0r9.execute-api.us-west-2.amazonaws.com/production/results/atrox/sync-dotenv)
[![Coverage Status](https://img.shields.io/codecov/c/github/atrox/sync-dotenv.svg?style=flat-square)](https://codecov.io/gh/Atrox/sync-dotenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/atrox/sync-dotenv?style=flat-square)](https://goreportcard.com/report/github.com/atrox/sync-dotenv)

Keep your `.env.example` in sync with changes to your `.env` file.

## Installation

### macOS (homebrew)
```sh
brew install atrox/tap/sync-dotenv
```

### Windows (scoop)
```ps
scoop bucket add sync-dotenv https://github.com/atrox/scoop-bucket
scoop install sync-dotenv
```

### Manually

Download the pre-compiled binaries from the [releases page](https://github.com/atrox/sync-dotenv/releases) and copy to the desired location.


## Usage

By default, `sync-dotenv` looks for a `.env` in your working directory and attempt to sync with .env.example
(or creates on, if the file does not exist) when no argument is provided.

If the flag `--watch` is provided, `sync-dotenv` will watch for changes and automatically update your example file.

## CLI

```sh
sync-dotenv helps you to keep your .env.example in sync with your .env file

Usage:
  sync-dotenv [flags]

Flags:
      --base string      base path for all paths (default ".")
      --env string       path to your env file (default ".env")
      --example string   path to your example env file (default ".env.example")
  -h, --help             help for sync-dotenv
  -v, --verbose          enable verbose messages
      --version          version for sync-dotenv
  -w, --watch            watch for file changes and update the example file automatically
```

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/atrox/sync-dotenv/issues)
- Fix bugs and [submit pull requests](https://github.com/atrox/sync-dotenv/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features
