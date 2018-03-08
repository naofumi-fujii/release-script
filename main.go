package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	branchName := fmt.Sprintf("%s", os.Args[1])

	// git branch -D $(BRANCH_NAME)
	out, err := exec.Command("git", "branch", "-D", branchName).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// git checkout develop
	out, err = exec.Command("git", "checkout", "develop").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// git pull
	// git checkout -b $(BRANCH_NAME)
	// git push --set-upstream origin $(BRANCH_NAME)
	// git checkout master
	// git read-tree -u --reset $(BRANCH_NAME)
	// git commit -m '$(COMMIT_MSG)'
	// git push

}
