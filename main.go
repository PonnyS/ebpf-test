//go:build linux
// +build linux

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	objs := bpfObjects{}
	if err := loadBpfObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %s", err)
	}
	defer objs.Close()

	if err := objs.bpfPrograms.EbpfTest.Pin("/sys/fs/bpf/ebpf-test"); err != nil {
		log.Fatalf("pin objects: %s", err)
	}
	defer func() {
		log.Println("unpin")
		objs.bpfPrograms.EbpfTest.Unpin()
	}()

	log.Println("Counting packets...")

	quit := make(chan struct{})
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				var value uint64
				if err := objs.PktCount.Lookup(uint32(0), &value); err != nil {
					log.Fatalf("reading map: %v", err)
				}
				log.Printf("number of packets: %d\n", value)
			case <-quit:
				log.Println("quit")
				return
			}
		}
	}()

	<-ch
	quit <- struct{}{}
	time.Sleep(2 * time.Second)
}
