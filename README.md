# sync-dotenv

[![Release](https://img.shields.io/github/release/atrox/sync-dotenv.svg?style=flat-square)](https://github.com/atrox/sync-dotenv/releases/latest)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fatrox%2Fsync-dotenv%2Fbadge&style=flat-square)](https://actions-badge.atrox.dev/atrox/sync-dotenv/goto)
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

### Linux (deb/rpm)

Download the `.deb` or `.rpm` from the [releases page](https://github.com/atrox/sync-dotenv/releases) and install with `dpkg -i` and `rpm -i` respectively.

### Manually

Download the pre-compiled binaries from the [releases page](https://github.com/atrox/sync-dotenv/releases) and copy to the desired location.


## Usage

`sync-dotenv`, by default, looks for a `.env` file in your working directory and synchronizes those keys with your `.env.example` file.  Values will **not** be synchronized but existing ones in your example file will be kept.

If your files have different names, you can use the flags `--env .secrets` and `--example .secrets.example`.

If the flag `--watch` is provided, `sync-dotenv` will watch for changes in your working directory and automatically updates your example file.

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

## Jetbrains IDE

<img align="right" height="200" src="/assets/jetbrains.png?raw=true">

You can automatically call `sync-dotenv` via [File Watchers](https://plugins.jetbrains.com/plugin/7177-file-watchers). Enable the plugin, import the [following XML](/assets/watchers.xml) and your `.env.example` now automatically updates if you change your `.env` file.
<br><br>

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/atrox/sync-dotenv/issues)
- Fix bugs and [submit pull requests](https://github.com/atrox/sync-dotenv/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features
