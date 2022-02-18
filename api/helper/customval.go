package helper

func ValidasiBalik(name string) string {
	switch name {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "oneof":
		return "Value can only"
	case "max":
		return "More character length"
	case "min":
		return "Less character length"
	default:
		return ""
	}

}
