package file_replacer

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"

	"github.com/justtrackio/structmd/cmd/config-structs-as-markdown/file_parser"
	"github.com/justtrackio/structmd/cmd/config-structs-as-markdown/markdown_generator"
)

const (
	BeginTag = "[structmd]:# ("
	EndTag   = "[structmd end]:#"
)

func ReplaceFile(fileName string) {
	lines := extractLines(fileName)

	for i := range lines {
		if strings.HasPrefix(lines[i], BeginTag) && strings.HasSuffix(lines[i], ")") {
			markdown := computeStructsMarkdown(lines[i][len("[structmd]:# (") : len(lines[i])-1])

			lines[i] = fmt.Sprintf("%s\n%s", lines[i], markdown)
		}
	}

	lines = removeOldStructsMarkdown(lines)

	writeLines(fileName, lines)
}

func removeOldStructsMarkdown(lines []string) []string {
	result := make([]string, 0, len(lines))
	insideOldMarkdown := false

	for i := 0; i < len(lines); i++ {
		switch {
		case strings.HasPrefix(lines[i], EndTag):
			insideOldMarkdown = false

			result = append(result, lines[i])
		case strings.HasPrefix(lines[i], BeginTag):
			insideOldMarkdown = true

			result = append(result, lines[i])
		case !insideOldMarkdown:
			result = append(result, lines[i])
		}
	}

	return result
}

func computeStructsMarkdown(line string) string {
	parts := strings.Split(line, " ")
	codeFile := parts[0]
	structNames := parts[1:]

	fSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fSet, codeFile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	structDefinitions := file_parser.ParseStructDefinitions(astFile, structNames)
	markdown := markdown_generator.Generate(structDefinitions...)

	return markdown
}

func extractLines(fileName string) []string {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(input), "\n")
}

func writeLines(filename string, lines []string) {
	output := strings.Join(lines, "\n")

	err := ioutil.WriteFile(filename, []byte(output), 0o644)
	if err != nil {
		panic(err)
	}
}
