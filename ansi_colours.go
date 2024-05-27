package ansi_colours

var colourMap = map[string]string{
	"reset":   "\033[0m",
	"red":     "\033[31",
	"green":   "\033[32",
	"yellow":  "\033[33",
	"blue":    "\033[34",
	"magenta": "\033[35",
	"cyan":    "\033[36",
	"gray":    "\033[37",
	"white":   "\033[97",
}

func assemble(colour string, bold bool, text string) string {
	r := colourMap[colour]

	if bold {
		r += ";1m"
	} else {
		r += "m"
	}

	r += text + colourMap["reset"]

	return r
}

func Red(text string) string {
	return assemble("red", false, text)
}

func Green(text string) string {
	return assemble("green", false, text)
}

func Yellow(text string) string {
	return assemble("yellow", false, text)
}

func Blue(text string) string {
	return assemble("blue", false, text)
}

func Magenta(text string) string {
	return assemble("magenta", false, text)
}

func Cyan(text string) string {
	return assemble("cyan", false, text)
}

func Gray(text string) string {
	return assemble("gray", false, text)
}

func White(text string) string {
	return assemble("white", false, text)
}

func BrightRed(text string) string {
	return assemble("red", true, text)
}

func BrightGreen(text string) string {
	return assemble("green", true, text)
}

func BrightYellow(text string) string {
	return assemble("yellow", true, text)
}

func BrightBlue(text string) string {
	return assemble("blue", true, text)
}

func BrightMagenta(text string) string {
	return assemble("magenta", true, text)
}

func BrightCyan(text string) string {
	return assemble("cyan", true, text)
}

func BrightGray(text string) string {
	return assemble("gray", true, text)
}

func BrightWhite(text string) string {
	return assemble("white", true, text)
}
