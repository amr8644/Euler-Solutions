package main

import "fmt"

//  Problem #2
//  Each new term in the Fibonacci sequence is generated by
//  adding the previous two terms. By starting with 1 and 2,
//  the first 10 terms will be:

//  1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

//  By considering the terms in the Fibonacci sequence whose values do not
//  exceed four million, find the sum of the even-valued terms.

//   Intialize the 1st 2nd and next term

func main()  {
    t_1 := 0
    t_2 := 1
    sum := 0
    nxt_t := t_1 + t_2

    for i := 1; i < 32; i++ {
        t_1 = t_2
        t_2 = nxt_t
        nxt_t = t_1 + t_2
        if nxt_t % 2 == 0 {
            sum = sum + nxt_t
        }
    }

    fmt.Println(sum)
}
