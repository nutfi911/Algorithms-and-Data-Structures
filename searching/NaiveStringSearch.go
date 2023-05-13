package searching

func NaiveStringSearch(mainStr string, subStr string) int {
	count := 0
	for _, char := range mainStr {
		for i, char2 := range subStr {
			if char != char2 {
				continue
			}
			if i == len(subStr)-1 {
				count++
			}
		}

	}

	return count
}
