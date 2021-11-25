package file_parser_test

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/justtrackio/structmd/domain"
	"github.com/justtrackio/structmd/file_parser"
	"github.com/stretchr/testify/assert"
)

func TestParseStructDefinitions(t *testing.T) {
	code := `package apiserver

// MySettingsStruct does smth.
// line 1
// line two
type MySettingsStruct struct {
	// Port does smth else.
	// Port related comment
	Port        string              ` + "`" + `cfg:"port" default:"8080"` + "`" + `
}
`

	fSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fSet, "", code, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	actual := file_parser.ParseStructDefinitions(astFile, []string{"MySettingsStruct"})
	assert.Len(t, actual, 1)

	expected := domain.StructDefinition{
		Name:         "MySettingsStruct",
		GodocComment: "MySettingsStruct does smth.\\nline 1\\nline two",
		Fields: []domain.StructField{
			{
				Name:         "Port",
				DefaultValue: "8080",
				Type:         "string",
				GodocComment: "Port does smth else.\\nPort related comment",
			},
		},
	}
	assert.Equal(t, expected, actual[0])
}
