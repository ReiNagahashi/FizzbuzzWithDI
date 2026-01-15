//go:build wireinject
// +build wireinject
package main    //  

import (
	"tdd/core"
	fizzbuzzsequenceprintergood "tdd/fizzbuzz_sequence_printer_good"

	"github.com/google/wire"
)

func InitializeFizzBuzzPrinter() (*fizzbuzzsequenceprintergood.FizzBuzzSequencePrinter, error){
	wire.Build(
		fizzbuzzsequenceprintergood.NewFizzBuzzSequencePrinter, //クライアントコードのインスタンス→コンバーターインスタンス、その結果を出直するためのインターフェースをメンバーとして持つ
		core.NewNumberConverter, // ルールのリストを持ったコンバーターインスタンスの生成
		fizzbuzzsequenceprintergood.NewConsoleOutput, //プリンターが出力するためのインターフェースを実装したインスタンスの生成
		fizzbuzzsequenceprintergood.ProvideRules, //所望の数だけルールを定義してリストとして作成する
	)

	return nil, nil

}

