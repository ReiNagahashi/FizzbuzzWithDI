package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCyclicNumberRule_Match(t *testing.T){
	rule := NewCyclicNumberRule(3, "Fizz")
	assert.Equal(t, true, rule.Match("", 3))
	assert.Equal(t, true, rule.Match("", 6))
	assert.Equal(t, false, rule.Match("", 1))

}

func TestApply(t *testing.T){
	rule := NewCyclicNumberRule(0, "Fizz")
	assert.Equal(t, "Fizz", rule.Apply("", 0))
	assert.Equal(t, "FizzFizz", rule.Apply("Fizz", 0))

}