package main

func Sum(vs []int) int {
	ans := 0
	for _, v := range vs {
		ans += v
	}
	return ans
}
func Keys(m map[any]any) []any {
	ans := make([]any, 0, len(m))
	for k := range m {
		ans = append(ans, k)
	}
	return ans
}
