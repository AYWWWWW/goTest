package main

var digitMap = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	results := []string{}
	if len(digits) < 1 {
		return []string{}
	}
	letters := digitMap[int(digits[0]-'0')]
	if len(digits) == 1 {
		for _, v := range letters {
			results = append(results, string(v))
		}
	} else {
		for _, v := range letters {
			for _, vs := range letterCombinations(string(digits[1:])) {
				results = append(results, string(v)+string(vs))
			}
		}
	}
	return results
}

func main() {
}
