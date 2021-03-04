package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func putNum(ch chan int){
	for i := 2;i < 50000; i++{
		ch <- i
	}
	//关闭ch
	close(ch)
}
//从ch中取出数据，判断是否为素数，如果是，放入primech
func primeNum(ch chan int, primech chan int, exitch chan bool) {
	var flag bool
	for {
		num, ok := <-ch
		if !ok {
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //该num不是素数
				flag = false
				break
			}
		}
		if flag {
			primech <- num
		}
		exitch <- true
	}
}

func main(){
	ch := make(chan int,1000)
	primech := make(chan int, 50000)
	exitch := make(chan bool, 50000)
	//开启1个协程，向通道中输入1000个数
	go putNum(ch)
	//开启8个协程
	for i := 0; i < 100; i++ {
		go primeNum(ch, primech, exitch)
	}
	//主线程
	go func(){
		for i := 0; i < 8; i++ {
			<-exitch
		}
	}()

	for {
		res, ok := <-primech
		if !ok{
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}
}


