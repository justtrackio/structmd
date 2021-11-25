## Tool to generate markdown documentation for structs

structmd is a cmd line app which takes in markdown files and adds to them documentation for all structures in certain golang source code files.

#### How to run
In your markdown file (i.e. README.md) you can add lines like the 2 lines below:

`[structmd]:# (pkg/apiserver/cors.go MySettingsStruct MyOtherStruct)`
`[structmd end]:#`

Notice on the first line the `[structmd]:#` keyword followed by a single source code file `pkg/apiserver/cors.go` then a list of struct names. The second line is also mandatory, as it marks the end of this current `structmd` block.

Then you run the structmd cmd line app: 
`./structmd README1.md README2.md`

This app takes in a list of markdown files, parses them, and replaces all `[structmd]:#` lines with a list of struct descriptions.

#### Output table
* title: struct name followed by its godoc description
* first column: the names of all public fields of a given struct
* second column: an explanation of that that field does, taken from its godoc comments
* third column: the default value for that field, taken from its tag

#### Implementation

The code creates the AST for a given file, then traverses it using ast.Inspect. The nature of the tree is that for a given node ast.Inspect will be behave like this:
```
Inspect traverses an AST in depth-first order: It starts by calling f(node); node must not be nil. 
If f returns true, Inspect invokes f recursively for each of the non-nil children of node, followed by a call of f(nil).
```

In practice, for a given node f() will be invoked like this: 
```
   f(node)
   for child := range node.children {
      f(child) 
      Inspect(child, f)
   }
   f(nil)
```

#### Example
For this input:
```package apiserver

// MySettingsStruct does smth.
// line 1
// line two
type MySettingsStruct struct {
	// Port does smth else.
	// Port related comment
	Port        string              `cfg:"port" default:"8080"`
}
```

Our tool will produce the following output:

##### Struct **MySettingsStruct**

MySettingsStruct does smth.\nline 1\nline two

| field       | type     | default     | description     |
| :------------- | :----------: | :----------: | -----------: |
| Port | string | 8080 | Port does smth else.\nPort related comment |

#### Notes
* the following godoc rule is being obeyed: if a struct or field has one or more lines of comments, the first comment beginning with the name of that struct/field is considered the beginning of its godoc comment
* the source file in the `[structmd]:#` line needs to be relative to the location of the repo's base directory, not relative to the location of the .md file
