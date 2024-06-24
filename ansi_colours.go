package ansi_colours

type open_close struct {
	open  string
	close string
}

var colourMap = map[string]open_close{
	"black":   open_close{"30", "0"},
	"red":     open_close{"31", "0"},
	"green":   open_close{"32", "0"},
	"yellow":  open_close{"33", "0"},
	"blue":    open_close{"34", "0"},
	"magenta": open_close{"35", "0"},
	"cyan":    open_close{"36", "0"},
	"white":   open_close{"37", "0"},

	"bg_black":   open_close{"40", "0"},
	"bg_red":     open_close{"41", "0"},
	"bg_green":   open_close{"42", "0"},
	"bg_yellow":  open_close{"43", "0"},
	"bg_blue":    open_close{"44", "0"},
	"bg_magenta": open_close{"45", "0"},
	"bg_cyan":    open_close{"46", "0"},
	"bg_gray":    open_close{"47", "0"},

	"bold":      open_close{"1", "22"},
	"italic":    open_close{"3", "23"},
	"underline": open_close{"4", "24"},
	"blink":     open_close{"5", "25"},
	"reverse":   open_close{"7", "27"},
}

func assemble(colour string, text string) string {
	r := colourMap[colour]

	return "\033[" + r.open + "m" + text + "\033[" + r.close + "m"
}

func Black(text string) string {
	return assemble("black", text)
}

func Red(text string) string {
	return assemble("red", text)
}

func Green(text string) string {
	return assemble("green", text)
}

func Yellow(text string) string {
	return assemble("yellow", text)
}

func Blue(text string) string {
	return assemble("blue", text)
}

func Magenta(text string) string {
	return assemble("magenta", text)
}

func Cyan(text string) string {
	return assemble("cyan", text)
}

func White(text string) string {
	return assemble("white", text)
}

func BgBlack(text string) string {
	return assemble("bg_black", text)
}

func BgRed(text string) string {
	return assemble("bg_red", text)
}

func BgGreen(text string) string {
	return assemble("bg_green", text)
}

func BgYellow(text string) string {
	return assemble("bg_yellow", text)
}

func BgBlue(text string) string {
	return assemble("bg_blue", text)
}

func BgMagenta(text string) string {
	return assemble("bg_magenta", text)
}

func BgCyan(text string) string {
	return assemble("bg_cyan", text)
}

func BgWhite(text string) string {
	return assemble("bg_white", text)
}

func Bold(text string) string {
	return assemble("bold", text)
}

func Italic(text string) string {
	return assemble("italic", text)
}

func Underline(text string) string {
	return assemble("underline", text)
}

func Blink(text string) string {
	return assemble("blink", text)
}

func Reverse(text string) string {
	return assemble("reverse", text)
}
