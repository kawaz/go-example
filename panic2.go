package main

func foo() {
	bar();
}

func bar() {
	panic("パニックしました");
}

func main() {
	foo();
}

// vim: ts=4
