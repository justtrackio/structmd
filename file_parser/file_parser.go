package file_parser

import (
	"go/ast"
	"strings"

	"github.com/justtrackio/structmd/domain"
	"github.com/thoas/go-funk"
)

func ParseStructDefinitions(astFile *ast.File, requiredStructs []string) []domain.StructDefinition {
	result := make([]domain.StructDefinition, 0)

	var astIdentStack []*ast.Ident
	var astGenDeclStack []*ast.GenDecl
	ast.Inspect(astFile, func(x ast.Node) bool {
		if n, ok := x.(*ast.GenDecl); ok {
			if n == nil {
				astGenDeclStack = astGenDeclStack[:len(astGenDeclStack)-1]
			} else {
				astGenDeclStack = append(astGenDeclStack, n)
			}

			return true
		}

		if n, ok := x.(*ast.Ident); ok {
			if n == nil {
				astIdentStack = astIdentStack[:len(astIdentStack)-1]
			} else {
				astIdentStack = append(astIdentStack, n)
			}

			return true
		}

		s, ok := x.(*ast.StructType)
		if !ok {
			return true
		}

		structName := computeStructName(astIdentStack)
		if structName == "" {
			return true
		}

		structComments := computeStructComments(structName, astGenDeclStack)
		fields := extractFields(s.Fields.List)

		if funk.ContainsString(requiredStructs, structName) {
			result = append(result, domain.StructDefinition{
				Name:         structName,
				GodocComment: structComments,
				Fields:       fields,
			})
		}

		return false
	})

	return result
}

func extractFields(fieldList []*ast.Field) []domain.StructField {
	result := make([]domain.StructField, 0)

	for _, field := range fieldList {
		if field != nil {
			if len(field.Names) > 0 {
				current := domain.StructField{
					Name: field.Names[0].Name,
				}

				if field.Doc != nil {
					comments := make([]string, 0, len(field.Doc.List))

					for _, c := range field.Doc.List {
						comments = append(comments, c.Text)
					}

					current.GodocComment = extractGoDocComment(field.Names[0].Name, comments)
				}

				if field.Tag != nil {
					current.DefaultValue = extractDefaultValue(field.Tag.Value)
				}

				if ident, ok := field.Type.(*ast.Ident); ok {
					current.Type = ident.Name
				}

				result = append(result, current)
			}
		}
	}

	return result
}

func extractDefaultValue(tag string) string {
	if len(tag) < 3 {
		return ""
	}

	nonOuterQuotes := tag[1 : len(tag)-1]
	parts := strings.Split(nonOuterQuotes, " ")
	for _, p := range parts {
		if strings.HasPrefix(p, "default:") {
			result := p[len("default:"):]

			if strings.HasPrefix(result, "\"") && strings.HasPrefix(result, "\"") {
				result = result[1 : len(result)-1]
			}

			return result
		}
	}

	return ""
}

func computeStructName(astIdentStack []*ast.Ident) string {
	if len(astIdentStack) > 0 {
		astIdentParent := astIdentStack[len(astIdentStack)-1]

		return astIdentParent.Name
	}

	return ""
}

func computeStructComments(structName string, astGenDeclStack []*ast.GenDecl) string {
	if len(astGenDeclStack) > 0 {
		astGenDeclParent := astGenDeclStack[len(astGenDeclStack)-1]
		if astGenDeclParent.Doc == nil {
			return ""
		}

		comments := make([]string, 0, len(astGenDeclParent.Doc.List))

		for _, c := range astGenDeclParent.Doc.List {
			comments = append(comments, c.Text)
		}

		return extractGoDocComment(structName, comments)
	}

	return ""
}

func extractGoDocComment(name string, comments []string) string {
	var result strings.Builder

	godocFound := false
	for i, c := range comments {
		if !godocFound && (strings.HasPrefix(c, "// "+name)) {
			godocFound = true
		}

		if godocFound {
			if strings.HasPrefix(c, "// ") {
				c = c[3:]
			}

			result.WriteString(c)

			if i < len(comments)-1 {
				result.WriteString("\\n")
			}
		}
	}

	return result.String()
}
