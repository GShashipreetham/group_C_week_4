package main

func findMaxnumberMinnumber(a []int) (int, int) {
	maxnumber := a[0]
	minnumber := a[0]
	//giving the condition
	for _, n := range a {
		if n > maxnumber {
			maxnumber = n
		}
		if n < minnumber {
			minnumber = n
		}
	}
	return maxnumber, minnumber
	//here we are returning the function

}
