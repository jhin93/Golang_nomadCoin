package person

type Person struct { // 대문자로 설정되어야 export 가능
	name string
	age  int
}

func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
}

func (p Person) Name() string {
	return p.name
}
