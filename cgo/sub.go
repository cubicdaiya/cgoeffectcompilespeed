package main

/*
int sub(int a, int b) {
    return a - b;
}
*/
import "C"

func sub(a, b int) int {
	return int(C.sub(C.int(a), C.int(b)))
}
