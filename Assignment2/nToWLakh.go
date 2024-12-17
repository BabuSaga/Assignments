package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number (0-100000): ")
	fmt.Scan(&n)
	fmt.Print("Number in words: ")
	if n == 0 {
		fmt.Print("zero")
	} else {
		fmt.Print(onetolakh(n))
	}
	fmt.Println()
}

func onetolakh(n int) string {
	if n <= 19 {
		return ones(n)
	} else if n < 100 {
		return tens((n/10)*10) + " " + ones(n%10)
	} else if n < 1000 {
		return ones(n/100) + " hundred "  + onetolakh(n%100)
	} else if n < 20000 {
		return ones(n/1000) + " thousand " + onetolakh(n%1000)
	} else if n >= 20000 {
		return tens(n/1000) + " thousand " + onetolakh(n%1000)
	} else if n == 100000 {
		return "one lakh"
	}
	return ""
}

func ones(n int) string {
	var i string
	if n == 0 {
		i = "zero"
	} else if n == 1 {
		i = "one"
	} else if n == 2 {
		i = "two"
	} else if n == 3 {
		i = "three"
	} else if n == 4 {
		i = "four"
	} else if n == 5 {
		i = "five"
	} else if n == 6 {
		i = "six"
	} else if n == 7 {
		i = "seven"
	} else if n == 8 {
		i = "eight"
	} else if n == 9 {
		i = "nine"
	} else if n == 10 {
		i = "ten"
	} else if n == 11 {
		i = "eleven"
	} else if n == 12 {
		i = "twelve"
	} else if n == 13 {
		i = "thirteen"
	} else if n == 14 {
		i = "fourteen"
	} else if n == 15 {
		i = "fifteen"
	} else if n == 16 {
		i = "sixteen"
	} else if n == 17 {
		i = "seventeen"
	} else if n == 18 {
		i = "eighteen"
	} else if n == 19 {
		i = "nineteen"
	}
	return i
}

func tens(n int) string {
	var i string
	if n == 20 {
		i = "twenty"
	} else if n > 20 && n < 30 {
		i = "twenty "
		i = i + ones(n - 20)
	} else if n == 30 {
		i = "thirty"
	} else if n > 30 && n < 40 {
		i = "thirty "
		i = i + ones(n - 30)
	} else if n == 40 {
		i = "forty"
	} else if n > 40 && n < 50 {
		i = "forty "
		i = i + ones(n - 40)
	} else if n == 50 {
		i = "fifty"
	} else if n > 50 && n < 60 {
		i = "fifty "
		i = i + ones(n - 50)
	} else if n == 60 {
		i = "sixty"
	} else if n > 60 && n < 70 {
		i = "sixty "
		i = i + ones(n - 60)
	} else if n == 70 {
		i = "seventy"
	} else if n > 70 && n < 80 {
		i = "seventy "
		i = i + ones(n - 70)
	} else if n == 80 {
		i = "eighty"
	} else if n > 80 && n < 90 {
		i = "eighty "
		i = i + ones(n - 80)
	} else if n == 90 {
		i = "ninety"
	} else if n > 90 && n < 100 {
		i = "ninety "
		i = i + ones(n - 90)
	} else if n == 100 {
		i = "one hundred"
	}

	return i
}
