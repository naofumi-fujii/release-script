package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	branchName := fmt.Sprintf("%s", os.Args[1])

	// git branch -D $(BRANCH_NAME)
	out, err := exec.Command("git", "branch", "-D", branchName).Output()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(string(out))

	// git checkout develop
	out, err = exec.Command("git", "checkout", "develop").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// git pull origin develop
	out, err = exec.Command("git", "pull").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// git checkout -b $(BRANCH_NAME)
	out, err = exec.Command("git", "checkout", "-b", branchName).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// git push --set-upstream origin $(BRANCH_NAME)
	out, err = exec.Command("git", "push", "--set-upstream", "origin", branchName).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// git checkout master
	out, err = exec.Command("git", "checkout", "master").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// git read-tree -u --reset $(BRANCH_NAME)
	out, err = exec.Command("git", "read-tree", "-u", "--reset", branchName).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	commitMessage := "release " + time.Now().Format("2006-01-02")
	// git commit -m '$(COMMIT_MSG)'
	out, err = exec.Command("git", "commit", "-m", commitMessage).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// git push
	out, err = exec.Command("git", "push").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
