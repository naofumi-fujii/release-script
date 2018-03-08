package main

import (
	"fmt"
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

	// out, err := exec.Command("git", "branch").Output()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(out))

	out, err := exec.Command("git", "checkout", "-b", "develop").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
