package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

func pullStaticHtml(local string, remote string, key string) {

	ssh := fmt.Sprintf("/usr/bin/ssh -i %s  ", key)
	// rsync -Pav -e "ssh -i $HOME/.ssh/someKey" username@hostname:/from/dir/ /to/dir/
	for {
		cmd := exec.Command("rsync", "-Pavz", "--rsh", ssh, remote, local)
		// cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Error(err)
			break
		}
		log.Printf("%s\n", stdoutStderr)
		time.Sleep(5 * time.Second)
	}
}
