# Owing

> [!NOTE]
> Learning project!!!

> Track what you're owing your codebase

A CLI tool to find and report TODO, FIXME, HACK and NOTE comments in your source code.
Built with zero dependencies using only Go standard library.

## Features

- 🔍 Find TODO, FIXME, HACK and NOTE comments across your codebase
- 🚀 Fast directory scanning
- 📝 Clean, readable reports
- 🎯 zero external dependencies (stdlib only)

## Installation

```bash
go install github.com/matteo-gildone/owing@latest
```

## Usage

### Basic usage

```bash
# Scan current directory
owing .

# Scan specific directory
owing ./src
```

## Examples

### Text output

```bash
Found 5 TODOs in 2 files
FIXME: 1   HACK: 1   NOTE: 1   TODO: 2   
nested-folder/test-file.js (4):
   1 [TODO] Add handle of edge cases
   6 [HACK] Temporary random ID generator (not collision-safe)
  11 [TODO] Replace with proper validation logic
  17 [NOTE] Only checking name for now

test-file.js (1):
   1 [FIXME] this is wrong
```

## Supported comment  formats

`owing` recognizes comments in most programming languages:

```
// TODO: your message
// FIXME: your message
// HACK: your message
// NOTE: your message

/* TODO: your message */
# TODO: your message
-- TODO: your message
```
