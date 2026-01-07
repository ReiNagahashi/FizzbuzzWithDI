package spec

import "strconv"

type PassThroughRule struct{
	
}

func NewPassThroughRule() *PassThroughRule{
	return &PassThroughRule{}
}

func (p *PassThroughRule) Apply(carry string, n int)string{
	return strconv.Itoa(n)
}

func (p *PassThroughRule) Match(carry string, n int)bool{
	return carry == ""
}