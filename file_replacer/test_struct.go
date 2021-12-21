package file_replacer

import "time"

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

type MyTimeStruct struct {
	// Read timeout is the maximum duration for reading the entire request, including the body.
	Read time.Duration `cfg:"read" default:"60s" validate:"min=1000000000"`
}
