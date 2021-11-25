package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/justtrackio/structmd/file_replacer"
)

const helpMessage = `
In your markdown file (i.e. README.md) you can add lines like the two lines below: 

[structmd]:# (pkg/apiserver/cors.go MySettingsStruct MyOtherStruct)
[structmd end]:#

Notice on the first line the [structmd]:# keyword followed by a single source code file pkg/apiserver/cors.go then a list of struct names. The second line is also mandatory, as it marks the end of this current structmd block.

Run the all like this: ./cmd/replace_markdown/main.go README1.md README2.md

structmd takes in a list of markdown files, parses them, and replaces all [structmd]:# lines with a list of struct descriptions.
`

func main() {
	helpFlag := flag.Bool("help", false, "display usage instructions")
	flag.Parse()

	if *helpFlag {
		fmt.Print(helpMessage)

		return
	}

	for markdownFile := 1; markdownFile < len(os.Args); markdownFile++ {
		if !strings.HasSuffix(os.Args[markdownFile], ".md") {
			fmt.Printf("skipping %s as it is not a .md file\n", os.Args[markdownFile])

			continue
		}

		file_replacer.ReplaceFile(os.Args[markdownFile])
	}
}
