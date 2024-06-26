package specifications

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

type GurseAdapter func(name string) string

func (g GurseAdapter) Curse(name string) (string, error) {
	return g(name), nil
}
