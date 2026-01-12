package main

import (
	fizzbuzzsequenceprintergood "tdd/fizzbuzz_sequence_printer_good"
)

func main(){
	// printer := fizzbuzzsequenceprinter.FizzBuzzSequencePrinter{}
	// printer.PrintRange(1, 100)
	factory := fizzbuzzsequenceprintergood.NewFizzbuzzFactory()
	printer := factory.Create()
	printer.PrintRange(1, 100)
}