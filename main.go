package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "hogehoge"
	// reflect.Typeは型の名前を表す。
	// reflect.TypeOfで取得可能
	t := reflect.TypeOf(s)
	fmt.Println("type:", t)

	// reflect.Valueはreflection値の共通インタフェース
	// reflect.ValueOfで取得可能
	v := reflect.ValueOf(s)
	fmt.Println("value:", v)

	// 実は、ValueからもType型を取得することができる
	t2 := v.Type()
	fmt.Println("type:", t2)

	// Kindは、typeの具体的な種類を表す。
	k := v.Kind()
	fmt.Println("kind:", k)

	if v.Kind() == reflect.String {
		fmt.Println("value:", v.String())
	}

	type MyString string
	var s2 MyString = "foo"
	v3 := reflect.ValueOf(s2)
	// Kindは、元の型の情報を返す
	fmt.Println("value:", v3.Kind())



}
