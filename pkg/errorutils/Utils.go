package errorutils

func Check(e error) {
	if e != nil {
		panic("Wow wow")
	}
}
