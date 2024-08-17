package ascii

import "bufio"

// Store the art in a map
func ArtContainer(scanner *bufio.Scanner, ArtContainer map[rune][]string, counter int) map[rune][]string {
	for i := ' '; i <= '~'; i++ {
		// Read lines of ASCII art for each character
		for scanner.Scan() {
			if counter == 1 {
				counter++
				continue
			}
			if (counter-1)%9 != 0 {
				ArtContainer[i] = append(ArtContainer[i], scanner.Text())
				counter++
			} else {
				counter++
				break
			}
		}
	}
	return ArtContainer
}
