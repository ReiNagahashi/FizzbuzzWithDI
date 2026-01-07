package fizzbuzzsequenceprintergood

// 😄詳細への結合を避け、固有のロジックとそのための設計だけを残した実装

// ①インターフェースをメンバ変数と持つことで、別パッケージの変更による影響を防ぐ
type FizzBuzzSequencePrinter struct{
	
}

func NewFizzBuzzSequencePrinter() *FizzBuzzSequencePrinter{
	return &FizzBuzzSequencePrinter{}
}