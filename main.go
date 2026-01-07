package main

import (
	fizzbuzzsequenceprinter "tdd/fizzbuzz_sequence_printer_bad"
)

func main(){
	printer := fizzbuzzsequenceprinter.FizzBuzzSequencePrinter{}
	printer.PrintRange(1, 100)
}