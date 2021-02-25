// forms.go
package main

import (
    "html/template"
    "net/http"
    "os/exec"
    "fmt"
    "bufio"
)

type PageVariables struct {
	Success      bool
	Result       string
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

            HomePageVars := PageVariables{
              Success: true,
              Result: scanner.Text(),
            }

            tmpl.Execute(w, HomePageVars)
        }
    })

    http.ListenAndServe(":8080", nil)
}
