package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



type MockReplaceRuler struct{
	mock.Mock
}

func (m *MockReplaceRuler) Match(carry string, n int)bool{
	args := m.Called(carry, n)
	return args.Bool(0)
}

func (m *MockReplaceRuler) Apply(carry string, n int) string{
	args := m.Called(carry, n)
	return args.String(0)
}

func TestReplaceWithEmptyRules(t *testing.T){
	fizzBuzz := NewNumberConverter([]ReplaceRuler{})
	assert.Equal(t, "", fizzBuzz.Convert(1))
}



func TestConvertWithSingleRule(t *testing.T){
	// テスト対象の生成→ここでモックを実物の代わりに注入する
	fizzBuzz := NewNumberConverter([]ReplaceRuler{
		createMockRule(1, "", true, "Replaced"),
	})
	
	assert.Equal(t, "Replaced", fizzBuzz.Convert(1))
}


func TestConvertCompositingRuleResults(t *testing.T){
	fizzBuzz := NewNumberConverter([]ReplaceRuler{
		createMockRule(1, "", true, "Fizz"),
		createMockRule(1, "Fizz", true, "Buzz"),
	})
	assert.Equal(t, "FizzBuzz", fizzBuzz.Convert(1))
}



func TestConvertSkippingUnmatchedRules(t *testing.T){
	fizzBuzz := NewNumberConverter([]ReplaceRuler{
		createMockRule(1, "", false, "Fizz"),
		createMockRule(1, "", false, "Buzz"),
		createMockRule(1, "", true, "1"),
	})

	assert.Equal(t, "1", fizzBuzz.Convert(1))
}


// func TestConvertWithUnmatchedFizzBuzzRulesAndConstantRule(t *testing.T){
// 	rules := []ReplaceRuler{
// 		// 1なので、Fizzの代わりに”"
// 		createMockRule(),
// 		// 1なので、Buzzの代わりに""
// 		createMockRule(1, ""),
// 		// 1なので、"1"
// 		createMockRule(1, "1"),
// 	}

// 	fizzBuzz := NewNumberConverter(rules)

// 	assert.Equal(t, "1", fizzBuzz.Convert(1))
// }


func createMockRule(expectedNumber int, expectedCarry string, matchResult bool, replacement string) ReplaceRuler{
	mockRule := new(MockReplaceRuler)
	mockRule.On("Match", expectedCarry, expectedNumber).Return(matchResult).Once()
	mockRule.On("Apply", expectedCarry, expectedNumber).Return(replacement)
	
	return mockRule
}
