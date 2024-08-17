package ascii

// process input string and handle newline characters
func Line(s string) []string {
	tablo := []string{}
	str := ""

	if len(s) <= 1 {
		tablo = append(tablo, s)
		return tablo
	}

	value := true

	for i := 0; i < len(s); i++ {
		if string(s[i]) != "\\" && string(s[i]) != "n" {
			value = false
		}
		if i+1 < len(s) && s[i] == '\\' && s[i+1] == 'n' {
			if value {
				tablo = append(tablo, "\n")
				i++

			} else if i <= len(s)-2 {
				if i < len(s)-2 && s[i+2] != '\\' && s[i-1] == 'n' {
					i = i + 1
				} else if i < len(s)-2 && s[i-1] != 'n' && s[i+2] != '\\' {
					tablo = append(tablo, str)
					str = ""
					i++
				} else if str != "" {
					tablo = append(tablo, str)
					tablo = append(tablo, "\n")
					str = ""
					i++
				} else {
					tablo = append(tablo, "\n")
					i++
				}
			}
		}else if s[i] < ' ' || s[i] > '~' {
			if s[i] == '\r' {
				tablo = append(tablo, str)
				// tablo = append(tablo, "\n")
				str = ""
				i++
			} else {
				i++
			}
		} else {
			str += string(s[i])
		}
		if i == len(s)-1 && str != "" {
			tablo = append(tablo, str)
		}
	}
	return tablo
}
