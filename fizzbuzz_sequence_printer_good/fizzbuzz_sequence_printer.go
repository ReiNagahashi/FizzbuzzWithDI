package fizzbuzzsequenceprintergood

import (
	"fmt"
)

// 😄詳細への結合を避け、固有のロジックとそのための設計だけを残した実装
// インスタンス、インターフェースをメンバ変数として持たせるようにした
// →この実装によって、単体テストでは、モックオブジェクトを使って
// 「依存を意図した通りにコールしていれば良い」というスタンスで実装できる


type NumberConverter interface{
	Convert(n int)string
}

// outputInterfaceを準備した意図：クライアントコードの「printによる出力」という
// テスト困難さを、「writeメソッドを呼ぶ」という行為に抽象化するため
// これによりモックを使って「Writeメソッドが "Fizz" という引数で呼ばれたか？」
// を確認するだけで済む!
type OutputInterface interface{
	Write(data string )
}

// ①インターフェースをメンバ変数と持つことで、別パッケージの変更による影響を防ぐ
type FizzBuzzSequencePrinter struct{
	fizzBuzz NumberConverter
	output OutputInterface
}

func NewFizzBuzzSequencePrinter(converter NumberConverter, output OutputInterface) *FizzBuzzSequencePrinter{
	return &FizzBuzzSequencePrinter{fizzBuzz: converter, output: output}
}

func (p *FizzBuzzSequencePrinter) PrintRange(begin, end int) {
	
	for i := begin; i <= end; i++{
		text := p.fizzBuzz.Convert(i)
		formattedText := fmt.Sprintf("%d %s\n", i, text)
		p.output.Write(formattedText)
	}
}