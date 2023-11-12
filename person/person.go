package person

import "fmt"

type Person struct { // 대문자로 설정되어야 export 가능
	name string
	age  int
}

func (p *Person) SetDetails(name string, age int) { // Person을 *Person으로 바꾸지 않으면, Go 언어가 main.go의 nico를 복사한 복사본(메소드 리시버의 p)을 수정한다. 하지만 원본(main.go의 nico)는 수정되지 않는다.
	p.name = name
	p.age = age
	fmt.Println("SeeDetails nico :", p) // 결과 : SeeDetails nico : &{nico 12}
	// 결과 이유 : SeeDetails nico : &{nico 12} 라고 나온 이유는 p가 대변하는 nico가 struct라서 그런 것. 만약 타 자료형(ex int) 였으면 0xc000012088 형태의 자료형이 출력
	// Go 언어에서 포인터 변수를 출력할 때 & 기호가 결과에 포함되는 것은 주로 구조체를 가리키는 포인터에 해당합니다. 이는 구조체 인스턴스의 주소와 그 내용을 함께 표시하기 위한 Go의 표준 출력 방식입니다
}
