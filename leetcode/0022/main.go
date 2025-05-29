package main

func generateParenthesis(n int) []string {
	result := []string{}

	var dfs func(str string, open, close int)
	dfs = func(str string, open, close int) {
		if open > n || close > open {
			return
		}

		if open == n && close == n {
			result = append(result, str)
		}

		dfs(str+"(", open+1, close)
		dfs(str+")", open, close+1)
	}
	dfs("", 0, 0)

	return result
}
