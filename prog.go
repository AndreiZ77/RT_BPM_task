package main

import s "strings"
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strconv"
)

func main() {

  var inputStr string

  intCh := make(chan int, 5)
  total := 0
  cnt :=0

  for {
    fmt.Scan(&inputStr)
    if len(inputStr) != 0 {
      go getSite(inputStr, intCh)
      cnt = <-intCh
      fmt.Println("Count for " + inputStr + ":", strconv.Itoa(cnt))
      total += cnt
      inputStr = ""
    } else {
      fmt.Println("Total:", total)
      return
    }
  }


}

func getSite(url string, ch chan<- int) {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }

  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

  defer resp.Body.Close()
  ch <- s.Count(string(body), "Go")
}
