package main

//Fail early, fail often, fail fataly
func check(e error) {
	if e != nil {
		panic(e)
	}
}