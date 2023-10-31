package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	ip := os.Args[1]
	if len(ip) == 0 {
		log.Fatal("ip required")
	}

	ipDash := strings.ReplaceAll(ip, ".", "-")

	str := fmt.Sprintf("-P 2222 root@%s:/buildkite/builds/buildkite-agent-docker-%s/coinbase/c3-rewarder/gen/go/coinbase/rewarder/rewarder.pb.go ./gen/go/coinbase/rewarder/", ip, ipDash)
	splitStr := strings.Split(str, " ")

	out, err := exec.Command("scp", splitStr...).Output()
	if err != nil {
		fmt.Println(splitStr)
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

}
