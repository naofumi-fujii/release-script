package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// git branch -D $(BRANCH_NAME)
	// git checkout develop
	// git pull
	// git checkout -b $(BRANCH_NAME)
	// git push --set-upstream origin $(BRANCH_NAME)
	// git checkout master
	// git read-tree -u --reset $(BRANCH_NAME)
	// git commit -m '$(COMMIT_MSG)'
	// git push

	branchName := fmt.Sprintf("%s", os.Args[1])
	out, err := exec.Command("git", "checkout", "-b", branchName).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
