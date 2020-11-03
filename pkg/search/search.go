package search

import (
	"context"
	"strings"
	"sync"
)

//Result for
type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

//All for
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	
	var result Result
	chanResult := make(chan []Result, len(files))
	wg := sync.WaitGroup{}
	for _, file := range files {
		var resultTotal []Result	
		wg.Add(1)
		go func(file string) {
			newData := strings.Split(file,"\n")
			//log.Print(data)
			//log.Print(newData)

			for i, stroka := range newData {
				//log.Print(stroka)
				if strings.Contains(stroka, phrase) == true {
					result.Phrase = phrase
					result.Line = stroka
					result.LineNum = int64(i + 1)
					result.ColNum = int64(strings.Index(stroka, phrase))
					resultTotal = append(resultTotal, result)
				}
				
			}
			chanResult <- resultTotal
			wg.Done()
		}(file)
	}
	wg.Wait()
	close(chanResult)
	return chanResult
}
