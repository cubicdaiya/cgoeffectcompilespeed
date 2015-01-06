package main

/*
int div(int a, int b) {
    return a / b;
}
*/
import "C"

func div(a, b int) int {
	return int(C.div(C.int(a), C.int(b)))
}
