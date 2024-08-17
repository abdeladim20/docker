package ascii

func MyPrinter(ArtContainer map[rune][]string, input string) string {
	ArtToWrite := ""
	linee := 0
	for linee < 8 {
		for _, value := range input {
			ArtToWrite += ((ArtContainer)[value][linee])
		}
		if linee != 7 {
			ArtToWrite += "\n"
		}
		linee++
	}
	// ArtToWrite += "\n"
	return ArtToWrite
}