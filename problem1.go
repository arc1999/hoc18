package main

import "fmt"


var i =2
func main() {
	var a int
	var b int
fmt.Scanf("%d",&a)
fmt.Scanf("%d",&b)
	rec(a,b)

}
func rec ( a,b int) int{
	if(a<=b) {
		isPrime(a);
		a=a+1;
		rec(a,b);
	}
	return 0

}

func isPrime(a int) int {
	if(a%i!=0 && a!=2 && i<=(a/2)) {
		i=i+1
		isPrime(a)
	} else if(i>(a/2) || a==2) {
		i=2
		fmt.Println(a)
		return 0
	} else {
		i=2
		return 0
	}
	i=2
	return 0

}