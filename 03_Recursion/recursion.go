package main

func countdown(current int) {
	println(current)
	if current <= 0 { // base case
		return
	} else { // recursive case
		countdown(current - 1)
	}
}
func main() {
	countdown(100)

}
