package main 

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for i, char range tab {
		if f(char)
		count ++
	}
	return count
}
