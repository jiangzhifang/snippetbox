package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Base("/foo/bar/*.baz.js"))
	fmt.Println(filepath.Base("/foo/bar/baz"))
	fmt.Println(filepath.Base("/foo/bar/baz/"))

}
