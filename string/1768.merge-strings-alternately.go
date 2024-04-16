package string

func mergeAlternately(word1 string, word2 string) string {
	array1, array2 := []byte(word1), []byte(word2)
	var res []byte
	index := 0
	for index < len(array1) && index < len(array2) {
		res = append(res, array1[index])
		res = append(res, array2[index])
		index++
	}
	if index < len(array1) {
		res = append(res, array1[index:]...)
	}
	if index < len(array2) {
		res = append(res, array2[index:]...)
	}
	return string(res)
}
