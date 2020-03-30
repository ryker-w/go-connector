package stomp

//import (
//	"fmt"
//	"github.com/lishimeng/go-connector/stomp"
//	"time"
//)
//
//var stop = make(chan byte)
//var count = 10
//var index = 0
//var sub *stomp.Subscription
//var dest = "alex.pp.om"
//var mode = stomp.Queue
//
//func main() {
//	fmt.Println("app start...")
//
//	connector := stomp.New()
//	fmt.Println("connect to broker")
//	err := connector.Connect()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("connect success")
//
//	fmt.Println("subscribe queue")
//	sub, err = connector.Subscriber(mode, dest, func(bytes []byte) {
//
//		fmt.Printf("recv msg: %s\n", string(bytes))
//		index++
//		if index >= count {
//
//			fmt.Printf("has received %d messages\n", index)
//			fmt.Println("receiver work done.")
//			stop <- 0x00
//		}
//	})
//
//	if err != nil {
//		fmt.Println("subscribe failed")
//		connector.Close()
//		return
//	}
//
//	sender := connector.Sender(mode, dest)
//	go send(sender)
//	<-stop
//	fmt.Println("one work done")
//	<-stop
//	fmt.Println("second work done")
//
//	fmt.Println("unsubscribe queue")
//	connector.Unsubscribe(sub)
//	fmt.Println("close connection")
//	connector.Close()
//
//}
//
//func send(sender *stomp.Sender) {
//	defer func() {
//		fmt.Println("sender work done.")
//		<-stop
//	}()
//	fmt.Println("start sender...")
//
//	for i := 0; i < count; i++ {
//		fmt.Printf("message spawn %d\n", i+1)
//		err := sender.SendText(fmt.Sprintf("message [%d]", i+1))
//		if err != nil {
//			fmt.Println(err)
//		}
//		time.Sleep(3200 * time.Millisecond)
//	}
//}
