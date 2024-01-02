package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/jhin93/Golang_nomadCoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template // One big variable which manage whole templates.

type homeData struct {
	PageTitle string              // html 파일에 공유되어야 하기에 대문자
	Blocks    []*blockchain.Block // 블록을 렌더링하기 위해 선언. Declare as uppercase for sharing
}

// 첫번째 인자는 ResponseWriter. 이는 유저에게 보내고 싶은 데이터를 적는 것.
// 두번쨰 인자는 request pointer. request를 복사하려는 게 아니기에 포인터로 사용
func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()} // 모든 블록체인의 블록 렌더링
	templates.ExecuteTemplate(rw, "home", data)                      // 'templates'변수에 'ExecuteTemplate'메소드를 적용 'home'이라는 페이지를 data 변수를 적용해서 실행한다.
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil) // GET 요청은 단순히 "add" 템플릿 실행
	case "POST":
		r.ParseForm()                                           // HTTP 요청의 폼 데이터를 파싱하는 것입니다. ParseForm을 호출하면, 요청에 포함된 URL 쿼리 문자열과 요청 본문(body)의 폼 데이터가 파싱되어 Request 객체의 Form 필드에 저장됩니다
		data := r.Form.Get("blockData")                         // Request 객체의 Form 필드 중 "blockData" 라는 input에서 온 데이터를 변수 'data'에 담음.
		blockchain.GetBlockchain().AddBlock(data)               // 입력된 'blockData' 인풋의 데이터를 블록으로 추가.
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect) // 위 로직 실행 후 홈페이지로 리다이렉트하는 로직
	}
}

func Start() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))     // pages에 있는  모든 파일을 load한 결과물을 route안에서 호출.
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml")) // '/partial'에 있는 파일도 불러오기 위해 모든 파일이 담긴 'templates' 변수를 업데이트.
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	// .Fatal은 os.Exit(1) 다음에 따라나오는 error를 Print()
	// os.Exit(1)은 프로그램이 error code 1으로 종료되는 것
	log.Fatal(http.ListenAndServe(port, nil)) // ListenAndServe는 error를 반환하는 함수. 만약 에러가 있다면 log.Fatal발동. 아니라면 http.ListenAndServe 유지
}
