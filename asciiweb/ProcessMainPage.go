package asciiweb

import (
	"net/http"
	"text/template"
)

func ProcessMainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	} else {
		data := ""
		tmp1, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tmp1.Execute(w, data)
	}
}
