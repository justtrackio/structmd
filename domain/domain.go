package domain

type StructDefinition struct {
	Name         string
	GodocComment string
	Fields       []StructField
}

type StructField struct {
	Name         string
	DefaultValue string
	Type         string
	GodocComment string
}
