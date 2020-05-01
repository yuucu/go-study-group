package chapter3

import (
	"errors"
)

type Dog struct{}

func (d Dog) Bark() string {
	return "わんわん"
}

type Cat struct{}

func (c Cat) Crow() string {
	return "にゃーにゃ"
}

// 課題3
// この関数の引数はxの型は不定です。
// 型がDogの場合はBow()を実行した結果
// Catの場合はCrowを実行した結果
// その他の場合はerrorを返却してください。
func Kadai3(x interface{}) (string, error) {
	switch x_class := x.(type) {
	case Dog:
		return x_class.Bark(), nil
	case Cat:
		// x.(Cat).Crow(), nil
		return x_class.Crow(), nil
	default:
		return "", errors.New("invalid type")
	}
}
