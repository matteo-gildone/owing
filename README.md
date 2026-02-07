# owing

> Track what you're owing your codebase

A CLI tool to find and report TODO, FIXME, HACK and NOTE comments in your source code.
Built with zero dependencies using only Go standard library.

## Features

- 🔍 Find TODO, FIXME, HACK and NOTE comments across your codebase
- 📊 Multiple output formats: JSON and HTML
- 🚀 Fast directory scanning with configurable filter
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

# Scan with specific comment type
owing --type TODO ./src

# Output as JSON
owing --format JSON ./src

# Generate HTML report
owing --format html ./src > debt-report.html
```
### Options

| Option | Type | Description | Default |
| ------ | ---- | ----------- | ------- |
| `type` | string | Fileter by comment type: TODO, FIXME, HACK, NOTE | all |
| `format` | string | Output format: json, html | json |
| `exclude` | string | Comma separated list of directories to exclude | .git, vendor, node_modules |

## Examples

### JSON output

```
$ owing --format json ./src

[
  {
    "file": "./src/parser.go",
    "line": 42,
    "type": "TODO",
    "message": "Improve error message"
  },
  {
    "file": "./src/main.go",
    "line": 15,
    "type": "FIXME",
    "message": "Handle edge for empty files"
  }
]
```

### HTML report

```bash
$ owing --format html ./src > debt-report.html 
```
Opens a formatted HTML page with searchable, sortable table of all comments.

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
