package spec

type CyclicNumberRule struct{
	Base int
	Replacement string
}

func NewCyclicNumberRule(n int, replacement string) *CyclicNumberRule{
	return &CyclicNumberRule{Base:n, Replacement: replacement}
}


func (c *CyclicNumberRule) Match(carry string, n int) bool{
	return n % c.Base == 0
}

func (c *CyclicNumberRule) Apply(carry string, n int) string{
	return carry + c.Replacement
}