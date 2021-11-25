package markdown_generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justtrackio/structmd/cmd/config-structs-as-markdown/domain"
	"github.com/justtrackio/structmd/cmd/config-structs-as-markdown/markdown_generator"
)

func TestGenerate(t *testing.T) {
	input := domain.StructDefinition{
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

	actual := markdown_generator.Generate(input)
	expected := `##### Struct **MySettingsStruct**

MySettingsStruct does smth.\nline 1\nline two

| field       | type     | default     | description     |
| :------------- | :----------: | :----------: | -----------: |
| Port | string | 8080 | Port does smth else.\nPort related comment |
`

	assert.Equal(t, expected, actual)
}
