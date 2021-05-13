package main

import (
	"log"

	"github.com/tal-tech/go-zero/core/mr"
)

type DemoStruct struct {
	a, b, c, d int
}

func main() {
	demoStruct := &DemoStruct{}
	//mr.Finish 完成并发处理，全部完成或报错返回err变量
	err := mr.Finish(func() error {
		demoStruct.a = 1
		log.Println("demoStruct.a = 1")
		return nil
	}, func() error {
		demoStruct.b = 2
		log.Println("demoStruct.b = 2")
		return nil
	}, func() error {
		//time.Sleep(time.Second * 3)
		demoStruct.c = 3
		log.Println("demoStruct.c = 3")
		return nil
	}, func() error {
		demoStruct.d = 4
		log.Println("demoStruct.d = 4")
		return nil
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("%v\n", demoStruct)
	//sfdsfdsfdsf
	//////////////////////////////
	// s/dfsdfsdfsdfdsf
	//map reduce
	mapReduce, err := mr.MapReduce(func(source chan<- interface{}) {
		//初始化
		source <- demoStruct.a
		source <- demoStruct.b
		source <- demoStruct.c
		source <- demoStruct.d
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		//处理、计算
		i := item.(int)
		writer.Write(i * i)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		//组织返回结构
		var its []int
		for a := range pipe {
			x := a.(int)
			its = append(its, x)
		}
		writer.Write(its)
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	// mapReduce === var its []int
	log.Printf("%v\n", mapReduce)
	//ssssssssssssssssssssssssssssssssssssssssssssss
	///xxxxx
	// ...sdfasdfasdfsadfsadfsadfsdsadfmain()
}
