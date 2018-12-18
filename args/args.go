package args

func Map(args []string) map[string]string {
	var mapper string
	result := make(map[string]string)

	for _, arg := range args[1:] {
		if arg[0] == '-' {
			mapper = arg[1:]
			result[mapper] = ""
		} else {
			if result[mapper] != "" {
				result[mapper] += " "
			}
			result[mapper] += arg
		}
	}

	return result
}