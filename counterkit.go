package kits2

type CounterKit struct {
	data   *Counter
	readme string //  注释
}

func NewCounterKit(readme string) *CounterKit {
	c := &CounterKit{}
	c.data = NewCounter()
	c.readme = readme
	return c
}

func (c *CounterKit) Show() string {
	str := c.title()
	str += c.Str()
	str += "\n"
	return str
}

func (c *CounterKit) title() string {
	str := ""
	str += "\n----------------------\n" + " \n计数器信息:" + c.readme + "\n"
	return str
}

func (c *CounterKit) Inc() {
	c.data.Add(1)
}
func (c *CounterKit) IncBy(num int64) {
	c.data.Add(num)
}
func (c *CounterKit) Get() int64 {
	return c.data.Get()
}

func (c *CounterKit) Str() string {
	return c.data.Str()
}
