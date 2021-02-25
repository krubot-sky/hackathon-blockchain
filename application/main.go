// forms.go
package main

import (
    "html/template"
    "net/http"
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

        HomePageVars := PageVariables{
          Success: true,
          Result: search,
        }

        tmpl.Execute(w, HomePageVars)
    })

    http.ListenAndServe(":8080", nil)
}
