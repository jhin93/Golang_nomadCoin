# Golang_nomadCoin


range 사용예제
```
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

data format 예제
```
import "fmt"

func main() {
	x := 412412
	fmt.Printf("%b\n", x) // binary
	fmt.Printf("%o\n", x) // octal
	fmt.Printf("%x\n", x) // hexadecimal
	fmt.Printf("%U\n", x) // hexadecimal
}
```

Sprintf 
```
func main() {
	x := 412412
	xAsBinary := fmt.Sprintf("%b\n", x) // binary로 format 된 x의 string
	fmt.Println(x, xAsBinary)
}
```

for loop 예제  
```
	foods := [3]string{"2", "e", "3"}
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
```

Slice : array의 무한대 버전. ex []string, []boolean

append(slice, elem1, elem2)