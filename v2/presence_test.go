package v2

var f = &Foo{
	ID:        1,
	FirstName: "Dark",
	LastName:  "Vador",
	Age:       30,
	Bar: &Bar{
		Name: "Test",
	},
}

var b = &Foo{
	ID:        2,
	FirstName: "Florent",
	LastName:  "Messa",
	Age:       28,
}
var c = &Foo{
	ID:        3,
	FirstName: "Harald",
	LastName:  "Nordgren",
	Age:       27,
}

var results = []*Foo{f, c}

type Person struct {
	name string
	age  int
}
