# Golang_nomadCoin

``` 
현재까지 들은 번호 : #8.0  
다음 들어야 하는 번호 : #8.1  
```

# Go standard 패키지 https://pkg.go.dev/std

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
// 54
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
package main

import (
    "fmt"
)

func main() {
	x := 412412
	xAsBinary := fmt.Sprintf("%b\n", x) // binary로 format 된 x의 string
	fmt.Println(x, xAsBinary) // 412412 1100100101011111100
	
	name := "John"
    age := 30

    // Using fmt.Sprintf to create a formatted string
    formattedString := fmt.Sprintf("My name is %s and I am %d years old.", name, age)

    fmt.Println(formattedString)
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

</br>
</br>
</br>
</br>

***append 메소드***
```go
slice := []int{1, 2, 3}  
slice = append(slice, 4, 5)  
// 이제 slice는 [1, 2, 3, 4, 5]입니다.  

firstSlice := []int{1, 2, 3}  
secondSlice := []int{4, 5, 6}  
combinedSlice := append(firstSlice, secondSlice...)  
// combinedSlice는 [1, 2, 3, 4, 5, 6]입니다.  
```

</br>
</br>
</br>
</br>

**Singleton Pattern**


싱글톤 패턴은 디자인 패턴 중 하나로, 특정 클래스의 인스턴스가 프로그램 전체에서 하나만 생성되도록 보장하는 패턴입니다. Go 언어에서의 싱글톤 패턴을 매우 간단하게 설명하면 다음과 같습니다:

1. 클래스 정의: Go에서는 클래스 대신 struct를 사용합니다. 싱글톤이 될 struct를 정의합니다.
2. 프라이빗 생성자: 다른 곳에서 이 struct의 인스턴스를 직접 생성하지 못하도록 생성자 함수를 프라이빗하게 만듭니다. Go에서는 생성자 대신 프라이빗한 초기화 함수를 사용할 수 있습니다.
3. 전역 인스턴스: 이 struct의 유일한 인스턴스를 저장할 전역 변수를 선언합니다.
4. 인스턴스 접근 함수: 이 전역 인스턴스에 접근할 수 있는 공개 함수를 제공합니다. 이 함수는 전역 인스턴스가 이미 생성되었는지 확인하고, 생성되지 않았다면 새로 생성하여 반환합니다.
```go
type MySingleton struct {
    // 필요한 필드들
}

var instance *MySingleton

func getInstance() *MySingleton {
    if instance == nil {
        instance = &MySingleton{
            // 초기화 코드
        }
    }
    return instance
}

// 여기서 getInstance 함수는 항상 동일한 MySingleton 인스턴스를 반환합니다. 이렇게 함으로써, 전체 프로그램에서 MySingleton의 단 하나의 인스턴스만 존재하게 되는 것입니다.
```

</br>
</br>
</br>
</br>

**Writer, fmt.Fprint**  
Go 프로그래밍 언어에서 "Writer"는 데이터를 쓸 수 있는 엔티티를 나타내는 인터페이스입니다. Go에서 인터페이스는 타입이 구현해야 하는 메서드 시그니처의 모음입니다. Writer 인터페이스는 io 패키지에 정의되어 있으며 가장 널리 사용되는 인터페이스 중 하나입니다.
</br>
</br>
간단히 말해, Writer 인터페이스는 데이터를 받아서 어딘가에 기록할 수 있는 모든 종류의 대상을 추상화합니다. 예를 들어, 파일, 메모리 버퍼, 네트워크 연결 등이 있을 수 있습니다. 여기서 fmt.Fprint 함수는 Writer 인터페이스를 만족하는 어떤 객체에도 데이터를 기록할 수 있습니다. 따라서 fmt.Fprint는 콘솔이 아닌, Writer 인터페이스를 구현하는 다양한 대상에 데이터를 출력할 수 있습니다.
</br>
</br>

```go
func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from home!") // Fprint는 콘솔이 아닌 Writer에 출력하는 것. data를 format해서 Writer에 보내는 것
}

```
http.ResponseWriter (rw): 이 인터페이스는 서버가 HTTP 클라이언트 (브라우저나 다른 HTTP 클라이언트)에 응답을 보내는 데 사용됩니다. ResponseWriter를 통해 HTTP 응답의 본문(body), 상태 코드(status code), 헤더(headers) 등을 설정할 수 있습니다.
</br>
</br>
*http.Request (r): 이 포인터는 현재 요청에 대한 모든 정보를 포함하고 있습니다. 클라이언트로부터 받은 데이터, URL, 헤더, 메소드(GET, POST 등) 등의 요청 정보를 담고 있습니다.
</br>
</br>
fmt.Fprint 함수는 첫 번째 인자로 Writer 인터페이스를 받고, 두 번째 인자부터는 출력하고자 하는 데이터를 받습니다. 이 경우 rw (즉, http.ResponseWriter)는 Writer 인터페이스를 구현하고 있기 때문에, fmt.Fprint를 사용하여 "Hello from home"이라는 문자열을 클라이언트에게 보낼 수 있습니다.
</br>
</br>
home 함수의 동작 구조는 다음과 같습니다:

1. 클라이언트(브라우저나 다른 HTTP 클라이언트)가 서버에 HTTP 요청을 보냅니다.
2. 서버는 home 함수를 호출하면서 ResponseWriter와 *http.Request를 인자로 넘깁니다.
3. home 함수는 fmt.Fprint를 사용해 rw (ResponseWriter)에 "Hello from home" 문자열을 기록합니다.
4. 이 문자열은 HTTP 응답의 일부로 클라이언트에게 전송되어, 클라이언트는 이 문자열을 받아 볼 수 있습니다.


</br>
</br>
</br>
</br>

**HTTP**  
```go
func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)

    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // 요청 메서드 (GET, POST, 등) 확인
    fmt.Fprintf(w, "Request method: %s\n", r.Method)

    // 요청 URL의 쿼리 파라미터 접근
    name := r.URL.Query().Get("name")
    if name != "" {
        fmt.Fprintf(w, "Hello, %s!\n", name)
    } else {
        fmt.Fprintf(w, "Welcome to the home page!\n")
    }
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    // 요청 헤더 정보 접근
    userAgent := r.Header.Get("User-Agent")
    fmt.Fprintf(w, "Your User-Agent is: %s\n", userAgent)
}

```
이 예시에서 homeHandler 함수는 URL의 쿼리 파라미터를 통해 이름을 받아 인사말을 반환합니다. 예를 들어, 사용자가 / 경로에 ?name=John 쿼리 파라미터를 포함하여 요청하면, 서버는 "Hello, John!"이라고 응답합니다.  
반면, aboutHandler 함수는 요청의 User-Agent 헤더를 읽어 클라이언트의 정보를 반환합니다. 이를 통해 클라이언트가 사용하는 브라우저나 운영 체제에 대한 정보를 얻을 수 있습니다.  
이처럼 *http.Request 객체를 사용하면 요청과 관련된 상세한 정보에 접근하고, 이를 바탕으로 동적인 응답을 생성할 수 있습니다.  

- *http.Request에 포인터(*)가 붙은 이유
1. 메모리 효율성: http.Request 구조체는 요청 URL, 헤더, 바디, 쿠키 등 다양하고 상대적으로 큰 양의 데이터를 포함할 수 있습니다. 포인터를 사용함으로써, 이러한 큰 데이터 구조체의 복사본을 만들지 않고, 원본 구조체에 대한 참조만을 전달합니다. 이는 메모리 사용량을 줄이고 성능을 향상시킵니다.
2. 변경 가능성: 함수 내에서 http.Request의 내용을 변경해야 할 경우, 포인터를 통해 원본 구조체에 직접 접근하고 수정할 수 있습니다. 값으로 전달할 경우, 수정 사항이 원본 구조체에 반영되지 않습니다.
3. 표준 라이브러리 설계: Go의 표준 라이브러리, 특히 net/http 패키지는 HTTP 요청과 응답을 처리하기 위해 이러한 포인터 기반의 설계를 사용합니다. 이는 Go 커뮤니티에서 널리 받아들여진 관행이며, 일관된 API 사용을 가능하게 합니다.

따라서, *http.Request에서 포인터를 사용하는 것은 메모리 관리 및 구조체의 변경 가능성을 고려한 설계 결정이며, Go 언어의 표준 라이브러리와 일관성을 유지하기 위한 것입니다.

</br>
</br>
</br>
</br>

**Fprintf**
```go
func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	// Marshal()은 Go의 데이터를 json으로 변환. json과 에러를 반환. JSON을 Go의 데이터로 바꿀 땐 Unmarshal
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Fprintf(rw, "%s", b) // rw에 string화("%s")된 b를 작성한다.
}
```
console이 아닌 writer에 작성하는 메소드.

</br>
</br>
</br>
</br>

**Sprintf**
fmt.Sprintf의 결과는 문자열이며, 이 문자열은 변수에 저장되거나 다른 함수에 전달될 수 있습니다. 이 함수는 실제로 값을 출력하지 않으며, 단지 포맷된 문자열을 반환합니다.
```go
name := "홍길동"
age := 30

formattedString := fmt.Sprintf("%s는 %d살입니다.", name, age)
// 결과: "홍길동는 30살입니다."
```
*서식 지정자의 예시
%v: 값을 기본 형식으로 출력
%d: 정수를 10진수 형식으로 출력
%f: 부동소수점 숫자를 출력
%s: 문자열 출력
%t: 불린 값(true 또는 false) 출력

</br>
</br>
</br>
</br>

**struct field tag**
```go
type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty는 해당 field의 value가 비어있을 경우, 해당 field 자체를 생략해준다.
}

```
struct field tag 작성방법. struct field를 소문자로 json으로 보여준다. 백틱(``) + 따옴표(")

</br>
</br>
</br>
</br>

**Interface**

Go 언어에서 interface는 특정 메서드 집합으로 정의된 타입입니다. 인터페이스는 구체적인 구현을 지정하지 않고, 특정 동작을 하는 메서드들의 시그니처(메서드 이름, 매개변수, 반환 타입)만을 정의합니다. 어떤 타입이 인터페이스에 정의된 모든 메서드를 구현하면, 그 타입은 해당 인터페이스를 구현한다고 간주합니다.

인터페이스의 주요 특징은 다음과 같습니다:

- 타입 추상화: 인터페이스는 구체적인 타입의 내부 구현을 숨기고, 해당 타입이 어떤 동작을 할 수 있는지만을 표현합니다.
- 다형성: 다양한 타입이 동일한 인터페이스를 구현할 수 있으며, 인터페이스 타입의 변수는 그 인터페이스를 구현하는 어떤 타입의 값이라도 저장할 수 있습니다. 이는 다양한 타입을 동일한 방식으로 처리할 수 있게 해줍니다.
- 계약의 정의: 인터페이스는 사용되는 메서드에 대한 계약을 정의합니다. 어떤 타입이 인터페이스를 구현하기 위해서는 인터페이스에 정의된 모든 메서드를 구현해야 합니다.


***interface를 직접 구현하는 경우***
```go
type Speaker interface {
    Speak() string
}
// Speaker라는 인터페이스는 Speak라는 메서드를 가지고 있습니다. 
// 어떤 타입이든 Speak() string 메서드를 구현하면 Speaker 인터페이스를 구현한 것으로 간주됩니다.

type Dog struct {}

func (d Dog) Speak() string {
    return "멍멍"
}

type Cat struct {}

func (c Cat) Speak() string {
    return "야옹"
}

func main() {
    var speaker Speaker
    speaker = Dog{}
    fmt.Println(speaker.Speak()) // 출력: 멍멍

    speaker = Cat{}
    fmt.Println(speaker.Speak()) // 출력: 야옹
}

```
위 예시에서 Dog와 Cat 타입은 모두 Speaker 인터페이스를 구현합니다. main 함수에서는 Speaker 타입의 변수 speaker를 사용하여 Dog와 Cat 객체의 Speak 메서드를 호출합니다. 이를 통해 다양한 타입을 동일한 인터페이스로 추상화하고 다형성을 실현할 수 있습니다.

***interface를 직접 구현하지 않는 경우***
```go
type URL string 

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}
```
위 예시의 'URL'이라는 type은 MarshalText()라는 golang의 내제된 interface를 사용함. 마치 .push()와 같다.
내재 interface 예시 - TextMarshaler(https://pkg.go.dev/encoding#TextMarshaler)

</br>
</br>
</br>
</br>

**JSON Encode, Decode**
```go
func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks()) // 뭔가를 http client에 write할때 Encoder를 만들어서 encode 함.
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody)) // Decoder(.NewDecoder)를 만들어서 request(r)의 body(Body)로부터 읽어온다(.Decode). 그리고 그 결과를 실제(&) addBlockBody에 담는다(&addblockBody)
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated) // 에러처리를 위해 header에 http status 작성. StatusCreated는 201을 의미
	}
}
```

</br>
</br>
</br>
</br>

**Gorilla**
링크 : https://gorilla.github.io/

사용을 위해선 'npm install 라이브러리' 처럼 라이브러리 설치를 해주어야 한다.
ex) 터미널에서 'go get -u github.com/gorilla/mux' 실행

Gorilla Toolkit은 Go 언어로 웹 애플리케이션을 개발하는 데 도움을 주는 모음집입니다. 이 툴킷은 여러 패키지로 구성되어 있으며 다양한 웹 개발 작업을 간편하게 수행할 수 있도록 도와줍니다. 아래는 Gorilla Toolkit의 몇 가지 중요한 패키지와 기능에 대한 간략한 설명입니다:

1. gorilla/mux: 이 패키지는 HTTP 요청 라우팅 및 URL 패턴 매칭을 위한 강력한 라우터입니다. 웹 애플리케이션의 다양한 엔드포인트를 정의하고 관리하는 데 사용됩니다.
ex) 변수를 받는 url 작성 - handler.HandleFunc("/blocks/{id}", blocks)

2. gorilla/sessions: 세션 관리를 위한 패키지로, 사용자 인증 및 상태 관리와 같은 기능을 구현할 때 유용합니다.

3. gorilla/websocket: 웹소켓을 지원하며, 실시간 통신을 구현하는 데 사용됩니다. 실시간 채팅, 게임, 알림 등을 개발하는 데 유용합니다.

4. gorilla/handlers: HTTP 핸들러 함수를 간편하게 구성할 수 있는 다양한 미들웨어 및 유틸리티 함수를 제공합니다.

5. gorilla/schema: 웹 폼 데이터를 Go 구조체로 파싱하고 검증하는 데 도움이 되는 패키지입니다.

</br>
</br>
</br>
</br>

**Adapter Pattern**
```go
func jsonContentTypeMiddleware(next http.Handler) http.Handler { // middleware. middleware는 마지막 목적지 전에 호출됨.
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) { // HandlerFunc는 타입, http.Handler는 interface. 근데 왜 리턴타입이 매칭이 안되는데 에러가 안날까?
		// http.HandlerFunc는 'adapter'라고 불리는데 이는 인자로 들어간 함수가 조건에 부합하면 필요한 함수를 구현해준다. 그래서 위 'type url' 사례처럼 직접 url 혹은 type을 만들지 않아도 adapter가 대신 만들어줌. 이를 adapter 패턴이라 한다.
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
```
</br>
</br>
</br>
</br>

***숫자:숫자***
```go
fmt.Println(os.Args[2:]) // [2:] 는 [2]부터 끝까지 를 의미한다. [2:5]는 [2]부터 [5]까지를 의미
```


</br>
</br>
</br>
</br>

***bolt***

https://github.com/boltdb/bolt
https://pkg.go.dev/github.com/boltdb/bolt#section-readme
GO의 key/value DB 라이브러리.