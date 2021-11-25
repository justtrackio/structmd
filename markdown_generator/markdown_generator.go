package markdown_generator

import (
	"fmt"
	"strings"

	"github.com/justtrackio/structmd/cmd/config-structs-as-markdown/domain"
)

func Generate(input ...domain.StructDefinition) string {
	var result strings.Builder

	for i, structDef := range input {
		result.WriteString(generateStruct(structDef))

		if i < len(input)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

func generateStruct(input domain.StructDefinition) string {
	var result strings.Builder

	result.WriteString(fmt.Sprintf("##### Struct **%s**\n\n%s\n\n", input.Name, input.GodocComment))
	result.WriteString("| field       | type     | default     | description     |\n")
	result.WriteString("| :------------- | :----------: | :----------: | -----------: |\n")

	for _, f := range input.Fields {
		result.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", f.Name, f.Type, f.DefaultValue, f.GodocComment))
	}

	return result.String()
}
