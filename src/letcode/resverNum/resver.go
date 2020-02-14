package resverNum


func isPalindrome(x int) bool {

	var number,result,tmp int

	number = x
	for number!=0 && number > 0 {
		tmp = number%10
		result = result*10+tmp
		number = number/10
	}

   if result == x {
   	return true
	}

	return false
}
