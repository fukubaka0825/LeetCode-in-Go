package _014_longest_common_prefix

import "strings"

/*Go Approach 1: Horizontal scanning*/
/*Time complexity : O(S), where S is the sum of all characters in all strings.*/
/*Space complexity : O(1). We only used constant extra space. */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// Assume prefix
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
