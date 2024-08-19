package main

import "fmt"

func main() {
	var input string
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&input)

	if isPalindrome(input) {
		fmt.Println("Kata tersebut adalah palindrome")
	} else {
		fmt.Println("Kata tersebut bukan palindrome")
	}
	main()
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
