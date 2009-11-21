package main

import (
	"fmt";
	"path";
	"os";
)

type SimpleVisitor struct{
}

func (v *SimpleVisitor) VisitDir(path string, d *os.Dir) bool {
	fmt.Printf("d: %s\n", path);
	return true;
}

func (v *SimpleVisitor) VisitFile(path string, d *os.Dir) {
	fmt.Printf("f: %s\n", path);
}

func main() {
	v := &SimpleVisitor{};
	errors := make(chan os.Error);
	path.Walk("/etc/httpd", v, errors);
}
