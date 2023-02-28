package conv_bool_string

func BoolToWord(word bool) string {
	if word == true {
		return "Yes"
	}
	return "No"
}
