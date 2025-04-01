package util
import (
	"time"
	"fmt"
)

func Sleep(i int){
	for  j:=i; j>0; j-- {
		fmt.Printf("\r冷却还剩 %d 秒", j)
		time.Sleep(time.Second)
	}
}