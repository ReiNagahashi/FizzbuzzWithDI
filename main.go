package main

import (
	"log"
)

func main(){
	// DIコンテナ&オートワイヤリング(対象のクラス群は手動で指定)
	printer, err := InitializeFizzBuzzPrinter()
	if err != nil{
		log.Fatal(err)
	}
	printer.PrintRange(1, 100)
}

