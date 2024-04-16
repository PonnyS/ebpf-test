# How to build
1. GOPACKAGE=main bpf2go -cc clang -cflags '-O2 -g -Wall -Werror' -target bpfel,bpfeb bpf main.bpf.c -- -I /ebpf-test/headers
2. CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -o main ./main.go ./bpf_bpfel.go
