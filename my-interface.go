package example_project

type MyInterface interface {
	MyMethod(a int, b string) (int, error)
	MySecondMethod(a int) int
}

type MyComponent struct {
	i MyInterface
}

func (c *MyComponent) MyMethod() (int, error) {
	return c.i.MyMethod(13, "Hello")
}

func (c *MyComponent) MySecondMethod() int {
	return c.i.MySecondMethod(1)
}

func NewMyComponent(i MyInterface) *MyComponent {
	return &MyComponent{i: i}
}
