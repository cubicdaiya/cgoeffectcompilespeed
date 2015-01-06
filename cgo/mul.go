package main

/*
int mul(int a, int b) {
    return a * b;
}
*/
import "C"

func mul(a, b int) int {
	return int(C.mul(C.int(a), C.int(b)))
}
