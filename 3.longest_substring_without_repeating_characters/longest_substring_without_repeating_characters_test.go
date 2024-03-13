package longest_substring_without_repeating_charactersc
import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	// Test case 1: No repeating characters
	input1 := "abcde"
	expected1 := 5
	if output1 := lengthOfLongestSubstring(input1); output1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, but got %d", expected1, output1)
	}

	// Test case 2: Repeating characters in the middle
	input2 := "abcabcbb"
	expected2 := 3
	if output2 := lengthOfLongestSubstring(input2); output2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, but got %d", expected2, output2)
	}

	// Test case 3: Repeating characters at the beginning
	input3 := "bbbbb"
	expected3 := 1
	if output3 := lengthOfLongestSubstring(input3); output3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, but got %d", expected3, output3)
	}

	// Test case 4: Repeating characters at the end
	input4 := "pwwkew"
	expected4 := 3
	if output4 := lengthOfLongestSubstring(input4); output4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, but got %d", expected4, output4)
	}

	// Test case 5: Empty string
	input5 := ""
	expected5 := 0
	if output5 := lengthOfLongestSubstring(input5); output5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, but got %d", expected5, output5)
	}
}