package ascii

// Write the art in a file
func ArtToWrite(fixedargument []string, ArtContainer map[rune][]string) string {
	art := ""
	for _, ono := range fixedargument {
		if ono == "\n" {
			art += "\n"
			continue
		} else {
			art += MyPrinter(ArtContainer, ono)
		}
	}
	return art
}
