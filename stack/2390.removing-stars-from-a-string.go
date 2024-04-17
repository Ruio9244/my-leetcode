package stack

func removeStars(s string) string {
	arr := []byte(s)
	var res []byte
	for i := 0; i < len(arr); i++ {
		if arr[i] != '*' {
			res = append(res, arr[i])
		} else {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		}
	}
	return string(res)
}
