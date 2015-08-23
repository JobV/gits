package gits

import (
	"fmt"
	"log"
	"os/exec"
)

// Clone runs git clone URL TARGET for you
func Clone(gitURL string, targetPath string) (err error) {
	cloneCmd := fmt.Sprintf("git clone %s", gitURL)
	cmd := exec.Command(cloneCmd)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	return
}
