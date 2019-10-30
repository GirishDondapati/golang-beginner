package main
/*

	A race condition occurs in Go when two or more goroutines try to access the same resource. It may happen when a variable
attempts to read and write the resource without any regard to other routines. For this case sol: We have to use Mutual Exclusion locks or sync.Mutex.

*/
import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

var wait sync.WaitGroup
var count int
func  increment(s string) {
	for i :=0;i<10;i++ {
		x := count
		x++;
		time.Sleep(time.Duration(rand.Intn(4))*time.Millisecond)
		count = x;
		fmt.Println(s, i,"Count: ",count)
	}
	wait.Done()
}

func main(){
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value " ,count)
}
