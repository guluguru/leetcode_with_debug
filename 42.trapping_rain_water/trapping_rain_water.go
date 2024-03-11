package trapping_rain_water

func trap(height []int) int {
	// 下标i初能接的雨水量由 min(leftMax[i], rightMax[i]) 决定
	// leftMax——从0到i-1的最大值
	// rightmax——从n-1到+1的最大值
	n := len(height)
	if n == 0 {
		return 0
	}
	res := 0
	l, r := 0, n-1
	leftMax, rightMax := height[0], height[n-1]

	for l <= r {
		leftMax = max(leftMax, height[l])
		rightMax = max(rightMax, height[r])
		if leftMax < rightMax {
			res += leftMax - height[l]
			l++
		} else {
			res += rightMax - height[r]
			r--
		}
	}
	return res
}
