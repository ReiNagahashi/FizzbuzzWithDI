package fizzbuzzsequenceprintergood

import (
	"fmt"
	"tdd/core"
	"tdd/spec"
)

// 「与えられたインスタンスを使うだけのロジック」から追い出された部分=生成
// を、依存性注入の考え方で補っていく
// 以下は、生成に関する知識を、できるだけ丁寧に切り分けて実装したもの

type FizzbuzzFactory struct{
	
}

func NewFizzbuzzFactory() *FizzbuzzFactory{
	return &FizzbuzzFactory{}
}

func (f *FizzbuzzFactory) Create() *FizzBuzzSequencePrinter{
	return NewFizzBuzzSequencePrinter(f.createFizzbuzz(), f.createOutput())
}

func (f *FizzbuzzFactory) createFizzbuzz() NumberConverter{
	return core.NewNumberConverter([]core.ReplaceRuler{
		f.createFizzRule(),
		f.createBuzzRule(),
		f.createPassThroughRule(),
	})
}

func (f *FizzbuzzFactory) createFizzRule() core.ReplaceRuler{
	return spec.NewCyclicNumberRule(3, "Fizz")
}

func (f *FizzbuzzFactory) createBuzzRule() core.ReplaceRuler{
	return spec.NewCyclicNumberRule(5, "Buzz")
}

func (f *FizzbuzzFactory) createPassThroughRule() core.ReplaceRuler{
	return spec.NewPassThroughRule()
}

func (f *FizzbuzzFactory) createOutput() OutputInterface{
	return NewConsoleOutput()
}

type ConsoleOutput struct{

}

func NewConsoleOutput() *ConsoleOutput{
	return &ConsoleOutput{}
}

func (co *ConsoleOutput) Write(data string){
	fmt.Println(data)
}