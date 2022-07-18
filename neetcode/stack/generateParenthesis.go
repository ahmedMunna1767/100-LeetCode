package stack

import "fmt"

func generateParenthesis(n int) []string {
	result := make([]string, 0)

	var genPan func(string, int, int)

	genPan = func(start string, open int, close int) {
		fmt.Println("Here", start)
		if open < close || open > 3 || close > 3 {
			return
		}
		if open == n && close == n {
			result = append(result, start)
			return
		} else {
			genPan(start+string("("), open+1, close)
			genPan(start+string(")"), open, close+1)
		}
	}

	genPan("", 0, 0)
	return result

}

func GenerateParenthesis() {
	fmt.Printf("%v", generateParenthesis(3))
}
