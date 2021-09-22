package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var mymap sync.Map

func wc(text *os.File, resultCh chan sync.Map, doneCh chan struct{}) {

	scanner := bufio.NewScanner(text)
	// value := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for i := range words {
			v, b := mymap.LoadOrStore(words[i], 0)
			if b {
				value := v.(int)
				value++
				// fmt.Println(words[i], " ", value)
				mymap.Store(words[i], value)
			} else {
				// fmt.Println(words[i], " ALMACENADO 0")
			}
		}
	}
	resultCh <- mymap
	doneCh <- struct{}{}
}

var mp = map[string]int{}

func main() {
	f, err := os.Create("mi_archivo.txt")
	if err != nil {
		panic(err)
	}

	var files [10]*os.File

	files[0], _ = os.Open("./text.txt")
	files[1], _ = os.Open("./text2.txt")
	files[2], _ = os.Open("./text3.txt")
	files[3], _ = os.Open("./text4.txt")
	files[4], _ = os.Open("./text5.txt")
	files[5], _ = os.Open("./text6.txt")
	files[6], _ = os.Open("./text7.txt")
	files[7], _ = os.Open("./text8.txt")
	files[8], _ = os.Open("./text9.txt")
	files[9], _ = os.Open("./text10.txt")

	outCh := make(chan sync.Map)
	doneWrite := make(chan struct{})

	go func() {
		for c := range outCh {
			fmt.Println("????")
			c.Range(func(key, value interface{}) bool {
				k := key.(string)
				v := value.(int)
				val, ok := mp[k]
				if ok {
					mp[k] = val + v
				} else {
					mp[k] = v
				}
				return true
			})
		}
		doneWrite <- struct{}{}
	}()

	doneCh := make(chan struct{})

	final := len(files) - 1

	for i := 0; i <= final; i++ {
		fmt.Printf("ejecutando %d\n", i)
		go wc(files[i], outCh, doneCh)
	}

	doneNum := 0
	for doneNum < 10 {
		<-doneCh
		fmt.Println("Termino uno")
		doneNum++
	}
	close(outCh)
	<-doneWrite
	fmt.Println("ListoA")
	for k, v := range mp {
		f.WriteString(k + " " + fmt.Sprint(v) + "\n")
	}

	fmt.Println("ListoB")

}
