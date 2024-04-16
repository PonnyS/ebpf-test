# How to build
1. `GOPACKAGE=main bpf2go -cc clang -cflags '-O2 -g -Wall -Werror' -target bpfel,bpfeb bpf main.bpf.c -- -I /ebpf-test/headers`
2. `CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -o main ./main.go ./bpf_bpfel.go`

# How to run
1. run: `./main`
2. add iptables rule: `iptables -t mangle -I POSTROUTING -m bpf --object-pinned /sys/fs/bpf/ebpf-test`

terminal print:
```
2024/04/16 11:37:52 number of packets: 0
2024/04/16 11:37:53 number of packets: 0
2024/04/16 11:37:54 number of packets: 0
2024/04/16 11:37:55 number of packets: 0
2024/04/16 11:37:56 number of packets: 0
2024/04/16 11:37:57 number of packets: 0
2024/04/16 11:37:58 number of packets: 0
2024/04/16 11:37:59 number of packets: 3042
2024/04/16 11:38:00 number of packets: 8872
2024/04/16 11:38:01 number of packets: 15375
2024/04/16 11:38:02 number of packets: 20769
2024/04/16 11:38:03 number of packets: 27420
^C2024/04/16 11:38:04 quit
2024/04/16 11:38:06 unpin
```
