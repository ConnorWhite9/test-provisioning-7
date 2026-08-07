package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "container-orchestration node agent starting...")
	// Node registration, Docker client initialization, heartbeat loop,
	// and gRPC connection to control plane will be wired here.
	select {}
}
