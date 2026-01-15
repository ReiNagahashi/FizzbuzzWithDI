package core

type ReplaceRuler interface{
	Match(carry string, n int)bool
	Apply(carry string, n int) string
}

type NumberConverter struct{
	Rules []ReplaceRuler
}

func NewNumberConverter(rules []ReplaceRuler)*NumberConverter{
	return &NumberConverter{Rules: rules}
}


func (c *NumberConverter) Convert(n int)string{
	carry := ""
	for _, rule := range c.Rules{
		if rule.Match(carry, n){
			carry = rule.Apply(carry, n)
		}
	}

	return carry
}


