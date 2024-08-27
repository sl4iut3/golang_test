package main

import ( 
  "fmt"
  "time"
)



func main(){
    var count int // Default 0

    cptr := &count

    go incr(cptr)

    time.Sleep(20*time.Second)

    fmt.Println("fin du main ",*cptr)
}

// Increments the value of count through pointer var
func incr(cptr *int) {
   //fmt.Println("debut de incr",*cptr)
    for i := 0; i < 1000; i++ {
            go func() {
                    *cptr = *cptr + 1
                    fmt.Println("ds incr",*cptr)
    time.Sleep(200*time.Millisecond)
            }()
      }
    }
