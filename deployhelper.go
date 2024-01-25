package main

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
	"os/exec"
)

const DevBranch = "dev"

func main() {
	//open current Git repository
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Printf("Error opening repository: %s\n", err)
		log.Fatal(err)
	}
	log.Printf("Opened current repository\n")

	// get the current branch
	branch, err := repo.Head()
	if err != nil {
		log.Printf("Error getting current branch: %s\n", err)
		log.Fatal(err)
	}

	currentBranchName := branch.Name().Short()
	log.Printf("Current branch: %s\n", currentBranchName)

	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
	})

	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			log.Printf("Already up to date\n")
		} else {
			log.Printf("Error pushing: %s\n", err)
			log.Fatal(err)
		}
	}

	log.Printf("Pushed %s branch\n", currentBranchName)

	// check out the dev branch
	worktree, err := repo.Worktree()
	if err != nil {
		log.Printf("Error getting worktree: %s\n", err)
		log.Fatal(err)
	}

	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(DevBranch),
		Force:  true,
	})

	if err != nil {
		log.Printf("Error checking out dev branch: %s\n", err)
		log.Fatal(err)
	}
	log.Printf("Checked out %s branch\n", DevBranch)

	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
	})

	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			log.Printf("Already up to date\n")
		} else {
			log.Printf("Error pulling: %s\n", err)
			log.Fatal(err)
		}
	}

	log.Printf("Pulled %s branch\n", DevBranch)

	//merge the current branch into the dev branch
	log.Printf("Merging %s into %s\n", currentBranchName, DevBranch)

	//since go git doesn't support merge yet, we'll use the command line
	var cmd *exec.Cmd = exec.Command("git", "merge", currentBranchName)

	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	err = cmd.Run()
	if err != nil {
		log.Printf("Error merging: %s\n", err)
		log.Fatal(err)
	}

	log.Printf("Merged %s branch to %s branch \n", currentBranchName, DevBranch)

	//push the dev branch
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
	})

	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			log.Printf("Already up to date\n")
		} else {
			log.Printf("Error pushing: %s\n", err)
			log.Fatal(err)
		}
	}
	log.Printf("Pushed %s branch\n", DevBranch)

	//check out the current branch
	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(currentBranchName),
		Force:  true,
	})

	if err != nil {
		log.Printf("Error checking out %s branch: %s\n", currentBranchName, err)
		log.Fatal(err)
	}
	log.Printf("Returned you to %s branch\n", currentBranchName)
}
