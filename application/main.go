// forms.go
package main

import (
    "html/template"
    "net/http"
    "os/exec"
    "fmt"
    "bufio"
    "strings"
)

type PageVariables struct {
  SuccessFalse bool
	Success      bool
	BookID       string
  BookName     string
  BookAuthor   string
}

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

              HomePageVars := PageVariables{
                Success: true,
                BookID: strings.TrimSuffix(strings.Split(payload,"\"")[16], "\\"),
                BookName: strings.TrimSuffix(strings.Split(payload,"\"")[4], "\\"),
                BookAuthor: strings.TrimSuffix(strings.Split(payload,"\"")[8], "\\"),
              }

              tmpl.Execute(w, HomePageVars)
            } else {
              HomePageVars := PageVariables{
                SuccessFalse: true,
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
