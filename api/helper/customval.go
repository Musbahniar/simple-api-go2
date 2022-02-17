package helper

func ValidasiBalik(name string) string {
	// message := fmt.Sprintf("Hello %v", name)
	// return message
	switch name {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "oneof":
		return "Pilih salah satu"
	default:
		return ""
	}

}
