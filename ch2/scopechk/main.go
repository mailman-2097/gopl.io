package main
import (
"os"
"log"
)
var cwd string
func main() {
	var err error
cwd, err = os.Getwd() // NOTE: wrong!
if err != nil {
log.Fatalf("os.Getwd failed: %v", err)
}
log.Printf("Working directory = %s", cwd)
}