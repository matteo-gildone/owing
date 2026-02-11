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

- [ ] Group comments by file
- [ ] Statistic summary
  - Total count
  - Count by type
  - Count by file
- [ ] Improve text format using grouping

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

- [ ] Color output
  - TODO: blue
  - FIXME: yellow
  - HACK: red
  - NOTE: green
- [ ] --no-color flag
- [ ] Respect NO_COLOR env var

---

## v2.0

**Goal** Sharable reports 

**Tasks**

- [ ] HTML output with basic template
- [ ] Group by file
- [ ] Basic CSS styling
- [ ] Generate with `owing --format html ./src > report.html`
