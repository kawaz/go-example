package main

import (
	//"flag";
	//"fmt";
	"io";
	"os";
	"tabwriter";
)

func usage() {
	println("Usage: " + os.Args[0]);
}

func main() {
	usage();
	tabw := tabwriter.NewWriter(os.Stdout, 1, 8, '\t', 0);
	io.Copy(tabw, os.Stdin);
}

// vim: ts=4
