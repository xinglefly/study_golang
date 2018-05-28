package test

import (
	"fmt"
	"unicode/utf8"
)

//寻找最长不含有重复字符的子串
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		//fmt.Println(i, ch)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("xywecwab"))


	/*map1 := map[string]string{
		"Java":    "java web",
		"Android": "app",
		"Golang":  "map",
		"Python":  "test",
	}
	fmt.Println(map1)

	for _, v := range map1 {
		fmt.Println(v)
	}

	var map2 map[string]int
	fmt.Println(map2)

	map3 := make(map[int]string)
	fmt.Println(map3)

	delete(map1, "Java")
	fmt.Println("Get map value:")
	if name, ok := map1["Java"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("no value")
	}*/

	s := "马丛爱你"
	fmt.Println(s, len(s))
	fmt.Printf("%X", s)

	for i, ch := range s {
		fmt.Println(i, ch)
	}

	r, size := utf8.DecodeRuneInString(s)

	fmt.Println(r, size)

	byte := []byte(s)
	for len(byte) > 0 {
		ch, sizes := utf8.DecodeRune(byte)
		byte = byte[sizes:]
		fmt.Printf("%c ",ch)
	}

	fmt.Println()

	for i,ch := range []rune(s)  {
		fmt.Printf("(%d, %c)", i,ch)
	}

}
