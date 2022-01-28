package captcha

// take variadic number of integer and find the sum
// of all digits that match the next digit in the list.
// The list is circular, so the digit after the last digit is the first digit in the list.
func SumNext(nums ...int) int {
	sumable := make([]int, 0)
	sum := 0

	for i, n := range nums {
		// if n == nums[i+1] then push n to sumable
		// if i == len(nums) - 1 then compare n == nums[0]
		next := i + 1
		if next == len(nums) {
			next = 0
		}

		if n == nums[next] {
			sumable = append(sumable, n)
		}
	}

	for _, n := range sumable {
		sum += n
	}

	return sum
}
