package asciiart

func BannerInput(args []string) (string, string) {
	banner := "standard"
	var input string
	switch len(args) {
	case 2:
		banner = args[1]
		input = args[0]
		if args[1] != "shadow" && args[1] != "standard" && args[1] != "thinkertoy" {
			InvalidBanner()
		}
	case 1:
		input = args[0]
	case 0:
		input = ""

	default:
		Usage()
	}

	return input, banner
}
