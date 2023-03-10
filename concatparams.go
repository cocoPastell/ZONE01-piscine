package piscine

func ConcatParams(args []string) string {
	var c string

	for i := 0; i < len(args); i++ {
		c = c + args[i]
		if i != len(args)-1 {
			c = c + "\n"
		}

	}
	return c
}
