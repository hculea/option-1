//go:build darwin && arm64
// +build darwin,arm64

package test

import _ "embed"

//go:embed option-1
var Hello string
