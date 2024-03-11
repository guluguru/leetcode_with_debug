package container_with_most_water

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max := 0
	for l < r {
		if height[l] < height[r] {
			max = maxMax(max, height[l], r-l)
			l++
		} else {
			max = maxMax(max, height[r], r-l)
			r--
		}
	}
	return max
}

func maxMax(max int, height int, width int) int {
	if max < height*width {
		return height * width
	}
	return max
}
