package main

import (
	"fmt";
)

type A struct {
	a int;
	b string;
}

func main() {
	//typeで定義された順番に初期値を代入する感じ
	var a1 A = A{1, "a1"};
	
	//フィールド名を指定した初期化も出来る
	var a2 A = A{a:2, b:"a2"};

	//フィールド順指定とフィールド名指定すうことは出来なくて、コンパイルエラーになる
	//var a3 A = A{3, b:"a3"};
	
	//ポインタの場合はこんな感じでOK
	var a4 *A = &A{4, "a4"};

	//newすることも出来る模様、newによって何が起こるのかの詳細はまだ理解していない(^^;
	var a5 *A = new(A);

	//varと型を書くのが面倒な場合は := を使う方が楽
	a6 := A{6, "a6"};
	a7 := &A{7, "a7"};

	//ポインタを初期化しないとnilが入る
	var a8 *A;

	//2次元配列の初期化とかはちょっとごちゃごちゃする感じですね(^^;
	a9 := [][]int{[]int{1,2}, []int{11,22,33}};
	
	//変数の内容確認には %#v が便利！
	fmt.Printf("%#v\n", a1);
	fmt.Printf("%#v\n", a2);
	//fmt.Printf("%#v\n", a3);
	fmt.Printf("%#v\n", a4);
	fmt.Printf("%#v\n", a5);
	fmt.Printf("%#v\n", a6);
	fmt.Printf("%#v\n", a7);
	fmt.Printf("%#v\n", a8);
	fmt.Printf("%#v\n", a9);
}

