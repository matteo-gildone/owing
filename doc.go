/*
Owing scans source code directories to find and report TODO, FIXME. HACK and NOTE comment.
It provides colored, grouped output with statistics to help developers track technical debt and outstanding work.

Installation

	go install github.com/matteo-gildone/owing@latest

# Usage

Scan a directory for TODO comments:

	owing ./src

The tool will output:
  - Total count of all TODOs found
  - Count by type (TODO, FIXED, HACK, NOTE)
  - TODOs grouped nby file with line numbers
  - Color-coded output (respects NO_COLOR)

# Example output

	 Found 5 TODOs in 2 files

	 FIXME: 1   HACK: 1   NOTE: 1   TODO: 2

		    nested-folder/test-file.js (4):
			 1 [TODO] Add handle of edge cases
			 6 [HACK] Temporary random ID generator (not collision-safe)
			11 [TODO] Replace with proper validation logic
			17 [NOTE] Only checking name for now

		    test-file.js (1):
			 1 [FIXME] this is wrong

# Supported comment types

Owing recognizes four types of comment:

  - TODO: Tasks to be completed
  - FIXME: Code that needs fixing
  - HACK: Temporary workaround or non-ideal solutions
  - NOTE: Important information or context

# Comment format

Comments must follow this format (with or without space after colon):

	// TODO: your message here
	# FIXME: your message here

# Accessibility

Owing respects terminal color preferences:

- Detects NO_COLORS environment variable
- Detects TERM=dumb for non-color terminals
- Output remains readable without colors

To disable colors:

	NO_COLORS owing ./src

# Zero dependencies

Owing was build using only the Go standard library with no external dependencies, making it lightweight
and easy to install.
*/
package main
