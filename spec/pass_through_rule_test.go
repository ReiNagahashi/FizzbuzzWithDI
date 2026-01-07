package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPassThroughRule_Apply(t *testing.T){
	rule := NewPassThroughRule()

	assert.Equal(t, "1", rule.Apply("", 1))
	assert.Equal(t, "2", rule.Apply("Fizz", 2))
}

func TestPassThroughRule_Match(t *testing.T){
	rule := NewPassThroughRule()

	assert.Equal(t, true, rule.Match("", 0))
	assert.Equal(t, false, rule.Match("Fizz", 0))
}
