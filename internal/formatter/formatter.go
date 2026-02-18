package formatter

import (
	"fmt"
	"html/template"
	"io"
	"sort"

	"github.com/matteo-gildone/owing/internal/reporter"
	"github.com/matteo-gildone/owing/internal/styles"
	"github.com/matteo-gildone/owing/internal/todo"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>TODO Report - {{.Total}} items</title>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&family=JetBrains+Mono:ital,wght@0,100..800;1,100..800&display=swap" rel="stylesheet">
    <style>
        /* Box sizing rules */
		*,
		*::before,
		*::after {
		  box-sizing: border-box;
		}
		
		/* Prevent font size inflation */
		html {
		  -moz-text-size-adjust: none;
		  -webkit-text-size-adjust: none;
		  text-size-adjust: none;
		}
		
		/* Remove default margin in favour of better control in authored CSS */
		body, h1, h2, h3, h4, p,
		figure, blockquote, dl, dd {
		  margin-block-end: 0;
		}
		
		/* Remove list styles on ul, ol elements with a list role, which suggests default styling will be removed */
		ul[role='list'],
		ol[role='list'] {
		  list-style: none;
		}
		
		/* Set core body defaults */
		body {
		  min-height: 100vh;
		  line-height: 1.5;
		}
		
		/* Set shorter line heights on headings and interactive elements */
		h1, h2, h3, h4,
		button, input, label {
		  line-height: 1.1;
		}
		
		/* Balance text wrapping on headings */
		h1, h2,
		h3, h4 {
		  text-wrap: balance;
		}
		
		/* A elements that don't have a class get default styles */
		a:not([class]) {
		  text-decoration-skip-ink: auto;
		  color: currentColor;
		}
		
		/* Make images easier to work with */
		img,
		picture {
		  max-width: 100%;
		  display: block;
		}
		
		/* Inherit fonts for inputs and buttons */
		input, button,
		textarea, select {
		  font-family: inherit;
		  font-size: inherit;
		}
		
		/* Make sure textareas without a rows attribute are not tiny */
		textarea:not([rows]) {
		  min-height: 10em;
		}
		
		/* Anything that has been anchored to should have extra scroll margin */
		:target {
		  scroll-margin-block: 5ex;
		}
       
        :root {
			--font-sans: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
			--font-mono: 'JetBrains Mono', 'SF Mono', Monaco, 'Cascadia Code', monospace;

            /* TODO Type Colors */
            --color-fixme: #f97316;
            --color-fixme-dark: #c2410c;
            --color-fixme-bg: #fff7ed;
            --color-fixme-border: #fed7aa;
           
            --color-todo: #3b82f6;
            --color-todo-dark: #1e40af;
            --color-todo-bg: #eff6ff;
            --color-todo-border: #bfdbfe;
           
            --color-hack: #ef4444;
            --color-hack-dark: #991b1b;
            --color-hack-bg: #fef2f2;
            --color-hack-border: #fecaca;
           
            --color-note: #22c55e;
            --color-note-dark: #166534;
            --color-note-bg: #f0fdf4;
            --color-note-border: #bbf7d0;
           
            /* Neutrals */
            --color-bg: #f5f5f5;
            --color-white: #ffffff;
            --color-text: #1a1a1a;
            --color-text-muted: #666666;
            --color-text-light: #999999;
            --color-border: #e5e5e5;
            --color-hover: #fafafa;
        }

		html, body {
		  height: 100%;
		  margin: 0;
		}
       
        body {
            font-family: var(--font-sans);
			font-optical-sizing: auto;
			font-weight: 400;
			font-style: normal;
            background: var(--color-white);
            color: var(--color-text);
            line-height: 1.6;
            padding: 0;
			margin: 0;
			min-height: 100vh;
		  	display: flex;
		  	flex-direction: column;
        }

		main {
		  flex: 1;
		}
       
        .container {
            max-width: 1400px;
            margin: 0 auto;
            background: var(--color-white);
            border-radius: 8px;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
        }
       
        header {
            padding: 1rem 2rem;
            border-bottom: 1px solid var(--color-border);
			background: var(--color-bg);
        }

		header h1 { margin: 0;}
       
        h1 {
			font-family: var(--font-mono);
            font-size: 1.75rem;
            font-weight: 700;
            color: var(--color-text);
            margin-bottom: 1.5rem;
        }
       
        .stats {
            display: flex;
            gap: 1rem;
            flex-wrap: wrap;
			margin-bottom: 2rem;
        }
       
        .stat-badge {
            display: inline-flex;
            align-items: center;
            gap: 0.5rem;
            padding: 0.5rem 1rem;
            border-radius: 6px;
            font-weight: 600;
            font-size: 0.9rem;
            border: 1px solid;
        }
       
        .stat-badge.FIXME {
            background: var(--color-fixme-bg);
            color: var(--color-fixme-dark);
            border-color: var(--color-fixme-border);
        }
       
        .stat-badge.TODO {
            background: var(--color-todo-bg);
            color: var(--color-todo-dark);
            border-color: var(--color-todo-border);
        }
       
        .stat-badge.HACK {
            background: var(--color-hack-bg);
            color: var(--color-hack-dark);
            border-color: var(--color-hack-border);
        }
       
        .stat-badge.NOTE {
            background: var(--color-note-bg);
            color: var(--color-note-dark);
            border-color: var(--color-note-border);
        }
       
        .stat-icon {
            width: 14px;
            height: 14px;
            border-radius: 2px;
        }
       
        .FIXME .stat-icon { background: var(--color-fixme); }
        .TODO .stat-icon { background: var(--color-todo); }
        .HACK .stat-icon { background: var(--color-hack); }
        .NOTE .stat-icon { background: var(--color-note); }
       
        main {
            padding: 2rem;
        }
       
        .file-section {
            margin-bottom: 2rem;
            border: 1px solid var(--color-border);
            border-radius: 8px;
            overflow: hidden;
        }
       
        .file-header {
			cursor: pointer;
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 1rem 1.5rem;
            background: var(--color-hover);
            border-bottom: 1px solid var(--color-border);
        }
       
        .file-name {
            font-weight: 600;
            color: var(--color-text);
            display: flex;
            align-items: center;
            gap: 0.5rem;
        }
       
        .file-icon {
            color: var(--color-text-muted);
        }
       
        .file-count {
            font-size: 0.85rem;
            color: var(--color-text-muted);
        }
       
        .todo-list {
            list-style: none;
        }
       
        .todo-item {
            display: flex;
            align-items: flex-start;
            gap: 1rem;
            padding: 1rem 1.5rem;
            border-bottom: 1px solid var(--color-bg);
        }
       
        .todo-item:last-child {
            border-bottom: none;
        }
       
        .todo-item:hover {
            background: var(--color-hover);
        }
       
        .todo-type {
            display: inline-flex;
            align-items: center;
            gap: 0.35rem;
            font-size: 0.75rem;
            font-weight: 600;
            padding: 0.25rem 0.6rem;
            border-radius: 4px;
            text-transform: uppercase;
            letter-spacing: 0.5px;
            white-space: nowrap;
        }
       
        .todo-type.FIXME {
            background: var(--color-fixme-bg);
            color: var(--color-fixme-dark);
        }
       
        .todo-type.TODO {
            background: var(--color-todo-bg);
            color: var(--color-todo-dark);
        }
       
        .todo-type.HACK {
            background: var(--color-hack-bg);
            color: var(--color-hack-dark);
        }
       
        .todo-type.NOTE {
            background: var(--color-note-bg);
            color: var(--color-note-dark);
        }
       
        .todo-type-icon {
            width: 6px;
            height: 6px;
            border-radius: 50%;
        }
       
        .FIXME .todo-type-icon { background: var(--color-fixme); }
        .TODO .todo-type-icon { background: var(--color-todo); }
        .HACK .todo-type-icon { background: var(--color-hack); }
        .NOTE .todo-type-icon { background: var(--color-note); }
       
        .todo-content {
            flex: 1;
            display: flex;
            align-items: baseline;
            gap: 0.75rem;
        }
       
        .line-number {
            font-family: var(--font-mono);
            font-size: 0.8rem;
            color: var(--color-text-light);
            min-width: 3rem;
        }
       
        .todo-message {
            color: var(--color-text);
            flex: 1;
        }
       
        footer {
            text-align: center;
            padding: 1.5rem;
            color: var(--color-text-muted);
            font-size: 0.85rem;
            border-top: 1px solid var(--color-border);
        }
       
        footer a {
            color: var(--color-todo);
            text-decoration: none;
        }
       
        footer a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
        <header>
            <h1>TODO Report</h1>
        </header>
       
        <main>
			<div class="stats">
				{{range $type, $count := .CountByType}}
				<div class="stat-badge {{$type}}">
					<span class="stat-icon"></span>
					<span>{{$type}} {{$count}}</span>
				</div>
				{{end}}
			</div>
            {{range $file, $todos := .GroupedByFile}}
            <details class="file-section">
                <summary aria-controls="todo-list-{{$file}}" id="todo-summary-{{$file}}" class="file-header">
                    <div class="file-name">
                        <span class="file-icon">📁</span>
                        {{$file}}
                    </div>
                    <span class="file-count">{{len $todos}} issues</span>
                </summary>
                <ul id="todo-list-{{$file}}" aria-labelledby="todo-summary-{{$file}}" class="todo-list">
                    {{range $todos}}
                    <li class="todo-item">
                        <span class="todo-type {{.Type}}">
                            <span class="todo-type-icon"></span>
                            {{.Type}}
                        </span>
                        <div class="todo-content">
                            <span class="line-number">#{{.Line}}</span>
                            <span class="todo-message">{{.Message}}</span>
                        </div>
                    </li>
                    {{end}}
                </ul>
            </details>
            {{end}}
        </main>
       
        <footer>
            Generated by <a href="https://github.com/matteo-gildone/owing">owing</a>
        </footer>
		<script>
			const details = document.querySelectorAll('details');
			details.forEach((el, index) => {
				const summary = el.querySelector('summary');
				if (index === 0) {
					el.open = true;
				}

				summary.setAttribute('aria-expanded', el.open);
				el.addEventListener('toggle', () => summary.setAttribute('aria-expanded', el.open));
			})

		</script>
</body>
</html>
`

func Text(w io.Writer, r reporter.Report) error {
	baseStyle := styles.NewStyles()
	fileStyle := baseStyle.Bold()
	dimStyle := baseStyle.Dim()
	fmt.Fprintf(w, "Found %d TODOs in %d files\n", r.Total, len(r.GroupedByFile))

	types := make([]string, 0, len(r.CountByType))

	for commentType := range r.CountByType {
		types = append(types, commentType)
	}

	sort.Strings(types)
	for _, commentType := range types {
		typeStyle := getStyleForType(commentType)
		fmt.Fprintf(w, "%s: %d   ", typeStyle.Render(commentType), r.CountByType[commentType])
	}

	fmt.Fprintln(w)

	files := make([]string, 0, len(r.GroupedByFile))

	for file := range r.GroupedByFile {
		files = append(files, file)
	}

	sort.Strings(files)

	for _, file := range files {
		todos := r.GroupedByFile[file]
		header := fmt.Sprintf("%s (%d):\n", file, len(todos))
		fmt.Fprint(w, fileStyle.Render(header))
		for _, t := range todos {
			typeStyle := getStyleForType(t.Type)
			fmt.Fprintf(w, "  %s %s %s\n", dimStyle.Render(fmt.Sprintf("%2d", t.Line)), typeStyle.Render(fmt.Sprintf("[%s]", t.Type)), t.Message)
		}
		fmt.Fprintln(w)
	}

	return nil
}

func Html(w io.Writer, r reporter.Report) error {
	tmpl := template.Must(template.New("report").Parse(htmlTemplate))
	return tmpl.Execute(w, r)
}

func getStyleForType(todoType string) styles.Style {
	base := styles.NewStyles()

	switch todoType {
	case todo.TypeTODO:
		return base.Cyan()
	case todo.TypeFIXME:
		return base.Yellow()
	case todo.TypeHACK:
		return base.Red()
	case todo.TypeNOTE:
		return base.Green()
	default:
		return base
	}
}
