package main

func makeParMap() map[string]string{
	parMap := make(map[string]string)
	parMap["}"] = "{"
	parMap["]"] = "["
	parMap[")"] = "("
	return parMap
}

func isValid(s string) bool {
	var stack []string
	for i := 0; i < len(s); i++{
		char := string(s[i])
		parMap := makeParMap()
		if _,ok := parMap[char]; ok{
			if len(stack) == 0{
				return false
			}

			n := len(stack)-1
			pop := stack[n]
			if parMap[char] != pop{
				return false
			}

			//pop
			stack = stack[:n]
			continue
		}
		//push
		stack = append(stack,char)
	}
	return len(stack) == 0
}
