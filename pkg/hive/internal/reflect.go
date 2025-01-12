// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package internal

import (
	"fmt"
	"path"
	"reflect"
	"regexp"
	"runtime"
)

var (
	baseNameRegex = regexp.MustCompile(`github.com/cilium/cilium/[\w\/]+/`)
)

func TrimName(name string) string {
	return string(baseNameRegex.ReplaceAll([]byte(name), []byte{}))
}

func PrettyType(x any) string {
	return TrimName(fmt.Sprintf("%T", x))
}

func FuncNameAndLocation(fn any) string {
	f := runtime.FuncForPC(reflect.ValueOf(fn).Pointer())
	file, line := f.FileLine(f.Entry())
	name := TrimName(f.Name())
	return fmt.Sprintf("%s (%s:%d)", name, path.Base(file), line)
}
