package asciiweb

import (
	"bufio"
	"net/http"
	"os"
	"text/template"

	"ascii-art-web/ascii"
)

func HandleInput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// handling the method
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	Input := r.FormValue("Input")
	Banner := r.FormValue("Banner")
	file, _ := os.Open(Banner + ".txt")
	if Banner != "standard" && Banner != "shadow" && Banner != "thinkertoy" {
		http.Error(w, "Banner Not Found", http.StatusBadRequest) // bad request ...
		return
	}

	// Defer closing the file until the function returns
	defer file.Close()

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	counter := 1
	// Process the input string and print the corresponding ASCII art
	fixedargument := ascii.NewLineHandler(Input)
	ArtContainer := make(map[rune][]string)
	Output := ascii.ArtToWrite(fixedargument, ascii.ArtContainer(scanner, ArtContainer, counter))

	tmp1, err := template.ParseFiles("ascii-art.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	tmp1.Execute(w, Output)
}
