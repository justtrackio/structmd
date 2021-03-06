package file_replacer_test

import (
	"io/ioutil"
	"testing"

	"github.com/justtrackio/structmd/file_replacer"
	"github.com/stretchr/testify/assert"
)

const (
	filename     = "README_TEST.md"
	fileContents = `[structmd]:# (test_struct.go MySettingsStruct MyOtherStruct MyTimeStruct)
[structmd end]:#
`
	expectedContents = `[structmd]:# (test_struct.go MySettingsStruct MyOtherStruct MyTimeStruct)
**MySettingsStruct**

MySettingsStruct does smth\nline 1\nline two

| field       | type     | default     | description     |
| :------------- | :----------: | :----------: | -----------: |
| Port | string | 8080 | Port does smth else\nPort related comment |

**MyOtherStruct**

MyOtherStruct does smth\nline one.\nline 2.

| field       | type     | default     | description     |
| :------------- | :----------: | :----------: | -----------: |
| Port | string | 2222 | Port does smth 11.\nPort related comment 2. |

**MyTimeStruct**



| field       | type     | default     | description     |
| :------------- | :----------: | :----------: | -----------: |
| Read | time.Duration | 60s | Read timeout is the maximum duration for reading the entire request, including the body. |

[structmd end]:#
`
)

func TestReplaceFile(t *testing.T) {
	err := ioutil.WriteFile(filename, []byte(fileContents), 0o644)
	assert.NoError(t, err)

	file_replacer.ReplaceFile(filename)

	contents, err := ioutil.ReadFile(filename)
	assert.NoError(t, err)

	actual := string(contents)
	assert.Equal(t, expectedContents, actual)

	err = ioutil.WriteFile(filename, []byte(fileContents), 0o644)
	assert.NoError(t, err)
}
