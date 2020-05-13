package main

import (
  "fmt"
  "time"
  "strconv"
)

func main() {
  c1 := make(chan []string) // init a string channel
  c2 := make(chan []string)

  go count("sheep", 500, c1) //begin go routine for function 'count'
  go count("fish", 1000, c2)

  for {
    select {
    case msg1 := <- c1:
      fmt.Println(msg1)
    case msg2 := <- c2:
      fmt.Println(msg2)
    }
  }

}

func count(thing string, waittime time.Duration, c chan []string) {

  for i := 1; i <= 5; i++ { // do it 5 times
    c <- []string{strconv.Itoa(i), thing} // send message to channel 'c'
    time.Sleep(time.Millisecond * waittime) //wait 500ms
  }
  close(c)
}
