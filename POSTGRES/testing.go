package main

import "fmt"

// func recurs (c []int32,counter int32,indexNext int) int32 {

//     // if (indexNext + 3) == len(c){
//     //     if (c[indexNext + 1] == 0 && c[indexNext+ 2] == 0) {
//     //         counter ++
//     //         indexNext = indexNext + 2
//     //     } else if (c[indexNext + 1] == 0 && c[indexNext + 2] == 1){
//     //         counter ++
//     //         indexNext ++
//     //     }
    
//     // } else if (indexNext + 2) == len(c) {
//     //     if (c[indexNext + 1] == 0 ){
//     //         counter ++
//     //         indexNext ++
//     //     }
//     // } 

//     if (indexNext + 1) == len(c) {
//         return counter 
//     } else {
//     	return recurs(c,counter,indexNext)
//     }

// } 

func main() {
	var data = []int32{0, 0, 1, 1, 0, 0}
	var counter int32 = 0
	var indexNext = 0

	fmt.Println(len(c))

	// counter = recurs(data,counter,indexNext)
	
	return counter

}
