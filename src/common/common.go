package common

func GetValue() (int, int) {
	return 10, 20
}

func GetDefault() int {
	return 0
}

func Sum(nums ...int) int64 {
	var s int64 = 0
	for _, v := range nums {
		s += int64(v)
	}
	return s
}

func Fact(n int) int {
	if n <= 1 {
		return 1
	}

	return n * Fact(n-1)
}
