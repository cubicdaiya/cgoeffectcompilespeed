package main

/*
int add(int a, int b) {
    return a + b;
}

int sub(int a, int b) {
    return a - b;
}

int mul(int a, int b) {
    return a * b;
}

int div(int a, int b) {
    return a / b;
}
*/
import "C"

func add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

func sub(a, b int) int {
	return int(C.sub(C.int(a), C.int(b)))
}

func mul(a, b int) int {
	return int(C.mul(C.int(a), C.int(b)))
}

func div(a, b int) int {
	return int(C.div(C.int(a), C.int(b)))
}
