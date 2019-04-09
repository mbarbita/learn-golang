package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	_ = iota
	AddOutDest
	ModOutDest
	DelOutDest
)

type Loop struct {
	id        int
	name      string
	desc      string
	sigMatrix [][]*IntSignal
}

type Out struct {
	dest chan int
	port int
}

type IntSignal struct {
	id     int
	name   string
	desc   string
	in     chan int
	outMap map[int]Out
}

func NewIntSignal(id int, kind, desc string) *IntSignal {
	return &IntSignal{
		id, //id
		fmt.Sprintf("%v %05v", kind, id), //name
		desc,              //desc
		make(chan int),    //in
		make(map[int]Out), //out
	}
}

func listSigComp(is *IntSignal) {
	fmt.Println("id  :", is.id)
	fmt.Println("name:", is.name)
	fmt.Println("desc:", is.desc)
	fmt.Println("in  :", is.in)
	for k, v := range is.outMap {
		fmt.Println("out :", k, v)
	}
}

func (isg *IntSignal) ModOut(cmd int, id int, ch chan int, port int) {
	switch cmd {
	case AddOutDest:
		out := &Out{ch, port}
		mutex.Lock()
		isg.outMap[id] = *out
		mutex.Unlock()
	case ModOutDest:
		out := &Out{ch, port}
		mutex.Lock()
		isg.outMap[id] = *out
		mutex.Unlock()
	case DelOutDest:
		mutex.Lock()
		delete(isg.outMap, id)
		close(ch)
		mutex.Unlock()
	}
}

func (isg *IntSignal) Start() {
	go func() {
		for {

			msg := <-isg.in
			// loop trough client map and send the message
			for _, v := range isg.outMap {
				v.dest <- msg
			}
			// log.Printf("sent message to %d dst", len(bcr.dst))
			// }
		}
	}()
}

func sendPulse(ch chan int, td time.Duration, msg int) {
	t := time.Tick(td)
	for _ = range t {
		ch <- msg
	}
}

func sendPulseMax(ch chan int, msg int) {
	for {
		ch <- msg
	}
}

func compFreq(isg *IntSignal, dstid int) {
	i := 1
	t := time.Tick(time.Second)
	for {
		select {
		case _, ok := <-isg.outMap[dstid].dest:
			if !ok {
				return
			}
			i++
		case <-t:
			// fmt.Printf("%10v imp/s\r", i)
			fmt.Printf("signal id: %v, signal name: %v, signal desc: %v, chan id: %v, chan: %v, port: %v, %10v imp/s\n",
				isg.id, isg.name, isg.desc, dstid, isg.outMap[dstid].dest, isg.outMap[dstid].port, i)
			i = 1
		}
	}
}

func compFreq2(isg *IntSignal, dstid int) {
	i := 1
	t := time.Tick(time.Second)
	for {
		select {
		case _, ok := <-isg.outMap[dstid].dest:
			if !ok {
				return
			}
			i++
		case <-t:
			// fmt.Printf("%10v imp/s\r", i)
			// fmt.Printf("signal id: %v, signal name: %v, signal desc: %v, chan id: %v, chan: %v, port: %v, %10v imp/s\n",
			// isg.id, isg.name, isg.desc, dstid, isg.outMap[dstid].dest, isg.outMap[dstid].port, i)
			i = 1
		}
	}
}

var loopMap = make(map[int]Loop)
var intSignals []*IntSignal
var intSignalsMap = make(map[int]*IntSignal)
var mutex = &sync.Mutex{}

func init() {

	// for i := 0; i < 3; i++ {
	// 	intSignals = append(intSignals, NewIntSignal())
	// 	intSignals[i].Start()
	// }

	for i := 1; i <= 2; i++ {
		intSignalsMap[i] = NewIntSignal(i, "PI", "TEST")
		// intSignalsMap[i].id = i
		// intSignalsMap[i].Name = fmt.Sprintf("AI %05v", i)
		intSignalsMap[i].Start()
	}

	// fmt.Println(intSignalsMap)
	// fmt.Println()
}

func main() {

	ch := make(chan int)
	intSignalsMap[1].ModOut(AddOutDest, 1, ch, 1)
	// listSigComp(intSignalsMap[1])
	go sendPulse(intSignalsMap[1].in, 100*time.Millisecond, 1)
	// go sendPulseMax(intSignalsMap[1].in, 1)
	go compFreq(intSignalsMap[1], 1)

	select {}
}
