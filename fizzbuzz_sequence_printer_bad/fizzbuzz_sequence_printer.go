package fizzbuzzsequenceprinterbad

import (
	"fmt"
	"tdd/core"
	"tdd/spec"
)

// 😩詳細に結合してしまっている実装

type FizzBuzzSequencePrinter struct{}

func NewFizzBuzzSequencePrinter() *FizzBuzzSequencePrinter{
	return &FizzBuzzSequencePrinter{}
}

func (p *FizzBuzzSequencePrinter) PrintRange(start, end int) {
	// ルールインスタンスを自力で定義・生成してる
	rules := []core.ReplaceRuler{
		spec.NewCyclicNumberRule(3, "Fizz"),
		spec.NewCyclicNumberRule(5, "Buzz"),
		spec.NewPassThroughRule(),
	}
	// NumberConbererインスタンスを自力で生成してる
	// →NumberCOnveterに変更が生じたら、このファイルに影響の可能性あり
	fizzBuzz := core.NewNumberConverter(rules)

	for i := start; i <= end; i++{
		// returnする代わりにprintによって外に結果を出してしまっている→単体テストが困難
		fmt.Println(fizzBuzz.Convert(i))
	}
}

