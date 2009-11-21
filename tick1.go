package main

import (
	"time";
	"fmt";
)

func main() {
	ticker := time.NewTicker(1000*1000*1000);
	for {
		ns := <-ticker.C;
		t := time.SecondsToLocalTime(ns /1000/1000/1000);
		fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", t.Year, t.Month, t.Day, t.Hour, t.Minute, t.Second);
	}
}
