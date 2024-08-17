package ascii

func NewLineHandler(s string) []string {
	var result []string
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' {
			if i+1 < len(s) && s[i+1] == '\n' {
				i++
			}
			if str != "" {
				result = append(result, str)
				str = ""
			}
			result = append(result, "\n")
		} else if s[i] == '\n' {
			if str != "" {
				result = append(result, str)
				str = ""
			}
			result = append(result, "\n")
		} else if s[i] >= ' ' && s[i] <= '~' {
			str += string(s[i])
		}
	}
	if str != "" {
		result = append(result, str)
	}
	return result
}
