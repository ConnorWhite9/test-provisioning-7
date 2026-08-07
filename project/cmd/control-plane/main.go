package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "container-orchestration control-plane starting...")
	// Initialization of state store, scheduler, controller manager,
	// REST API server, and gRPC API server will be wired here.
	select {}
}
