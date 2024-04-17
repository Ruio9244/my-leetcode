package array

func kidsWithCandies(candies []int, extraCandies int) []bool {
	length := len(candies)
	res := make([]bool, len(candies))
	if length == 0 {
		return res
	}
	maxValue := candies[0]
	for i := 1; i < length; i++ {
		if candies[i] > maxValue {
			maxValue = candies[i]
		}
	}
	for i := 0; i < length; i++ {
		if candies[i]+extraCandies >= maxValue {
			res[i] = true
		}
	}
	return res
}
