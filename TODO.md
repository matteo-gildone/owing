# owing development

## v1.0

**Goal:** Basic working tool

**Tasks**

- [x] Implement directory walker (fs.WalkDir)
  - [x] Skip .git, vendor, node_modules 
  - [x] Handle file read errors
- [x] Text output format `file:line [TYPE] message`
- [x] Make file for build/install
- [x] Tag and push v1.0.0

---

## v1.1

**Goal:** Better readability

**Tasks**

- [x] Group comments by file
- [x] Statistic summary
  - Total count
  - Count by type
  - Count by file
- [x] Improve text format using grouping

**Example output**
```bash
Found 5 TODOs in 2 files:

src/main.go (3):
  15: [TODO] Refactor this
  42: [FIXME] Handle error
  67: [TODO] Add validation
  
src/parser.go (2):
  12: [NOTE] Important context
  67: [TODO] Refactor this
```
---

## v1.2

**Goal** Prettier terminal output

**Tasks**

- [x] Color output
  - TODO: blue
  - FIXME: yellow
  - HACK: red
  - NOTE: green
- [x] Respect NO_COLOR env var

---

## v2.0

**Goal** Sharable reports 

**Tasks**

- [x] HTML output with basic template
- [x] Basic CSS styling
- [x] Generate with `owing --format html ./src > report.html`
