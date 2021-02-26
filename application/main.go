// forms.go
package main

import (
    //"encoding/json"
    "html/template"
    "net/http"
    "os/exec"
    "fmt"
    "bufio"
    "strings"
)

type PageVariables struct {
	Success      bool
	BookID       string
  BookName     string
  BookAuthor   string
}

//type Book struct {
//    BookID      string `json:"isbn"`
//    BookName    string `json:"title"`
//    BookAuthor  string `json:"author"`
//}

func main() {
    tmpl := template.Must(template.ParseFiles("index.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        search := r.FormValue("text")

        cmd := exec.Command("/usr/local/bin/peer","chaincode","invoke","-o","orderer:31010","-C","library","-n","bookstore","-c",`{"Args":["QueryBook","` + search + `"]}`)

        stderr, _ := cmd.StderrPipe()
        if err := cmd.Start(); err != nil {
            tmpl.Execute(w, nil)
            return
        }

        scanner := bufio.NewScanner(stderr)
        for scanner.Scan() {
            fmt.Println(scanner.Text())

            if strings.Contains(scanner.Text(), "payload") {
              payload := strings.Split(scanner.Text(),"payload:")[1]

              //book := Book {
              //  BookID: strings.Split(payload,"\"")[15],
              //  BookName: strings.Split(payload,"\"")[3],
              //  BookAuthor: strings.Split(payload,"\"")[7],
              //}

              //out, _ := json.Marshal(book)

              HomePageVars := PageVariables{
                Success: true,
                BookID: strings.Split(payload,"\"")[16],
                BookName: strings.Split(payload,"\"")[4],
                BookAuthor: strings.Split(payload,"\"")[8],
              }

              tmpl.Execute(w, HomePageVars)
            } else {
              HomePageVars := PageVariables{
                Success: false,
                BookID: "",
                BookName: "",
                BookAuthor: "",
              }

              tmpl.Execute(w, HomePageVars)
            }
        }
    })

    http.ListenAndServe(":8080", nil)
}
