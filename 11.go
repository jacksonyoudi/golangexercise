package main


func main() {
	j := 0
	for i := 1; j < 100; i++ {
		for j := 1; j < i; j++ {
			print('A')
		}
		print('\n')
	}
}
