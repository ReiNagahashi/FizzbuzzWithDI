package fizzbuzzsequenceprintergood

import (
	"fmt"
	"tdd/core"
	"tdd/spec"
)

// DIコンテナに登録するためのコンストラクタ(プロバイダと呼ばれる)を用意する。

// ルールの提供
// Wireは「[]ReplaceRuler」をどう作ればいいかを知らない→教えてあげる関数を作る
func ProvideRules() []core.ReplaceRuler{
	return []core.ReplaceRuler{
		spec.NewCyclicNumberRule(3, "Fizz"),
		spec.NewCyclicNumberRule(5, "Buzz"),
		spec.NewPassThroughRule(),
	}
}

// printerの提供
func NewFizzBuzzSequencePrinter(converter *core.NumberConverter, output OutputInterface) *FizzBuzzSequencePrinter{
	return &FizzBuzzSequencePrinter{
		converter: converter,
		output: output,
	}
}


// インターフェースを実装した具象クラスとして実装?
func NewConsoleOutput() OutputInterface{
	return &ConsoleOutput{}
}

type ConsoleOutput struct{
	
}

func (co *ConsoleOutput) Write(data string){
	fmt.Println(data)
}