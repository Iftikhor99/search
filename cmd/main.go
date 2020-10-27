package main

import (
	"sync"
	"strings"
	"time"
	"log"
	"context"
)

//Result for
type Result struct {
	Prase string
	Line string
	LineNum int64
	ColNum int64
}

func main() {



	//var files []string
	file1 := "1;+992000000001;10000000;food\n2;+992000000002;20000000;auto\n3;+992000000003;30000000;auto"
	file2 := "4;+992000000004;40000000;food\n5;+992000000005;50000000;food\n6;+992000000006;60000000;auto"
	files :=[]string{file1, file2}

	resultFound := all("food", files)
	log.Print(resultFound)
	// done := make(chan struct{})

	// go func (){
	// 	log.Print("goroutine")
	// }()
	// time.Sleep(time.Second*10)
	// done <- struct{}{}
	// ch := make(chan int)
	// root := context.Background()

	log.Print("Cancel")
	

		
	
	//log.Print(ctx.Err())
	//log.Print(ctx.Err() == context.Canceled) 
	//log.Print(root.Err())
	
}

func check(root context.Context) <- chan int {
	ch := make(chan int)

	
	ctx, cancel := context.WithCancel(root)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan int,ctx context.Context){
		for {
			select {
			case <- ctx.Done():
				log.Print("Done")
			defer	wg.Done()
					
				
				return
			case <- time.After(time.Second):
				log.Print("tick")
				ch <- 1
				
			}
			
		}	
	}(ch, ctx)

	go func() {
	close(ch)	
	wg.Wait	()
	}()
	
	cancel()
	

	log.Print("Cancel")
	
	return ch

}

func all(phrase string, files []string) []Result {
	var resultTotal []Result
	var result Result
	for _, file := range files {
		newData := strings.Split(file, "\n")
		//log.Print(data)
		//log.Print(newData)

		for i, stroka := range newData {
			//log.Print(stroka)
			if strings.Contains(stroka, phrase) == true {
				result.Prase = phrase
				result.Line = stroka
				result.LineNum = int64(i+1)
				result.ColNum = int64(strings.Index(stroka, phrase))
				resultTotal = append(resultTotal, result)
			}
			
	}	}
	return resultTotal
}