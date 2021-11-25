package file_replacer

// MySettingsStruct does smth
// line 1
// line two
type MySettingsStruct struct {
	// Port does smth else
	// Port related comment
	Port string `cfg:"port" default:"8080"`
}

// MyOtherStruct does smth
// line one.
// line 2.
type MyOtherStruct struct {
	// Port does smth 11.
	// Port related comment 2.
	Port string `cfg:"port" default:"2222"`
}
