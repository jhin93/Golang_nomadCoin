# Golang_nomadCoin


**range 사용예제**
```go
func plus(a ...int) int {
	total := 0
	for _, item := range a {
		total += item
	}
	return total
}

func main() {
	result := plus(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}

```

</br>
</br>
</br>
</br>
  
**data format 예제**
```go
import "fmt"

func main() {
	x := 412412
	fmt.Printf("%b\n", x) // binary
	fmt.Printf("%o\n", x) // octal
	fmt.Printf("%x\n", x) // hexadecimal
	fmt.Printf("%U\n", x) // hexadecimal
}
```
</br>
</br>
</br>
</br>

**Sprintf** 
```go
func main() {
	x := 412412
	xAsBinary := fmt.Sprintf("%b\n", x) // binary로 format 된 x의 string
	fmt.Println(x, xAsBinary) // 412412 1100100101011111100
}
```
</br>
</br>
</br>
</br>

**for loop 예제**  
```go
	foods := [3]string{"2", "e", "3"}
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
```
  
Slice : array의 무한대 버전. ex []string, []boolean
append 메소드 : .append(slice, elem1, elem2)
ex) slice = append(slice, "elem1")
  
</br>
</br>
</br>
</br>

**struct, method 사용예제**
```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() { // 'person' struct의 메소드
	fmt.Printf("Hello! my name is %s and i'm %d", p.name, p.age) // .Println으로 출력하면 '...name is %s and i'm %d' 라고 출력됨.
	// %s, %d 는 각각 string, decimal(10진수)를 의미함. 순서대로 %s에는 p.name, %d에는 p.age가 할당됨
}

func main() {
	nico := person{name: "nico", age: 12}
	nico.sayHello()
}
```
  
</br>
</br>
</br>
</br>

**Method Receiver**
```go
func (p *Person) SetDetails(name string, age int) { // 메서드 리시버 : (p *Person)
	p.name = name
	p.age = age
	fmt.Println("SeeDetails nico :", p)
}

```
Go 언어에서 SetDetails 함수 정의에서 나타나는 (p *Person) 부분을 "메서드 수신자(method receiver)"라고 부릅니다. 메서드 수신자는 해당 메서드가 연결될 구조체의 인스턴스를 정의합니다. 이 경우, SetDetails 메서드는 Person 구조체의 포인터 인스턴스(*Person)에 연결되어 있습니다.

Method Receiver는 메서드가 호출될 때 해당 구조체의 인스턴스에 대한 참조를 제공합니다. 이를 통해 해당 인스턴스의 필드를 읽거나 수정할 수 있습니다. 예를 들어, SetDetails 메서드 내에서 p.name과 p.age는 Person 인스턴스의 name과 age 필드를 참조하고 있습니다. 포인터(*Person)를 사용하는 것은 이 메서드가 호출될 때, 원본 Person 인스턴스를 직접 수정할 수 있도록 해줍니다.

이러한 특성 때문에, 구조체의 필드를 수정하는 메서드를 정의할 때는 일반적으로 구조체의 포인터(*Person)를 사용하는 것이 좋습니다. 이렇게 하면 메서드가 원본 구조체 인스턴스를 직접 수정할 수 있으므로, 변경사항이 해당 인스턴스에 반영됩니다.

*만약 메서드 리시버를 포인터로 설정하지 않을 경우
```go
// person/person.go
package person

import "fmt"

type Person struct { // 대문자로 설정되어야 export 가능
	name string
	age  int
}

func (p *Person) SetDetails(name string, age int) { // Person을 *Person으로 바꾸지 않으면, Go 언어가 main.go의 nico를 복사한 복사본(메소드 리시버의 p)을 수정한다. 하지만 원본(main.go의 nico)는 수정되지 않는다.
	p.name = name
	p.age = age
	fmt.Println("SeeDetails nico :", p)
}

// main.go
package main

import (
	"fmt"

	"github.com/jhin93/Golang_nomadCoin/person"
)

func main() {
	nico := person.Person{}             // nico 변수에 'Person' type 할당
	nico.SetDetails("nico", 12)         // 결과 : SeeDetails nico : {nico 12}
	fmt.Println("Main's nico : ", nico) // 결과 : Main's nico' { 0}
}

```
  
  

</br>
</br>
</br>
</br>

**구조체의 포인터 변수는 출력하면 '&구조체'의 형태로 출력됨**
```go
// main.go
package main

import (
	"fmt"

	"github.com/jhin93/Golang_nomadCoin/person"
)

func main() {
	nico := person.Person{}             // nico 변수에 'Person' type 할당
	nico.SetDetails("nico", 12)         // 결과 : SeeDetails nico : {nico 12}
	fmt.Println("Main's nico : ", nico) // 결과 : Main's nico' { 0}
}

// person/person.go
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
}
```
결과 이유 : SeeDetails nico : &{nico 12} 라고 나온 이유는 p가 대변하는 nico가 struct라서 그런 것. 만약 타 자료형(ex int) 였으면 0xc000012088 형태의 자료형이 출력
Go 언어에서 포인터 변수를 출력할 때 & 기호가 결과에 포함되는 것은 주로 구조체를 가리키는 포인터에 해당합니다. 이는 구조체 인스턴스의 주소와 그 내용을 함께 표시하기 위한 Go의 표준 출력 방식입니다.
</br>
</br>
</br>
</br>
**method receiver의 struct pointer instance가 '&'가 아니고 '*'인 이유**

직관적으로 생각하면 원본을 수정하기 위해 *가 아닌 메모리 주소 &가 메소드 리시버에 붙어야 한다고 보인다.  
그러나 실제로는 *를 해주어야 한다. 그래야 원본의 주소값 p(&{구조체} ex &{nico 12})에 직접적인 변형을 줄 수 있다.
메소드 리시버의 구조는 (1번 '인스턴스', 2번 '자료형 포인터 인스턴스') 이고, 1번이 주소이고 1번에 직접 변형을 가하는 2번이 포인터여야 하는 것이다. 

만약 주소형에 그냥 값을 할당하려고 하면 아래와 같은 에러가 나온다. 그래서 주소에 할당을 하려면 *b라고 *를 붙여줘야 하는 것
```go
func main() {
	a := 2
	b := &a
	b = 3 // cannot use 3 (untyped int constant) as *int value in
}
```
</br>
</br>
</br>
</br>

**One-way function**  
단방향 함수. 오직 한 방향으로만 작동하는 함수이다.  
결정론적이기에 일정한 입력값에 대해 일정한 결과값을 갖는다.  
ex) "test" = h_fn(x) => "awetawedgzderw"  
ex) "awetawedgzderw" = h_fn(x) => "test" ---> 이렇게 결과를 입력해도 입력값이 나오지 않는다. 단방향이기 때문.  