// execute git that tags a repository then confirms if the tagging was successfull

// You can use the Go os/exec package to execute git commands and check the result.
// Here's an example of how you can tag a Git repository using Go and confirm if the tagging was successful:

// This code assumes that you have Git installed and configured on your system.
// Replace "/path/to/your/repository" with the actual path to your Git repository and "v1.0.0" with the desired tag name.

// In this code:

//     We use os/exec to execute Git commands (git tag) from within the Go program.
//     First, we execute the git tag command to create a tag with the specified name.
//     Then, we execute the git tag -l command to list tags and verify if the tag was successfully created.
//     If the tag exists in the list of tags returned by git tag -l, we consider the tagging process successful.
//     If there are any errors during the execution of Git commands, the program exits with an error message.

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Define the repository path
	repoPath := "/path/to/your/repository"

	// Define the tag name
	tagName := "v1.0.0"

	/*
		Execute git tag command to create a tag
		script to tag and push plugin project:
		 	go mod tidy
			git add <go-filenam>
			git commit -m "set version v0.0.1"
			git tag v0.0.1
			git push origin v0.0.1
	*/
	cmd := exec.Command("git", "-C", repoPath, "tag", tagName)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error creating tag: %v", err)
	}

	// Execute git tag command again to verify if the tag was created
	cmd = exec.Command("git", "-C", repoPath, "tag", "-l", tagName)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error verifying tag: %v", err)
	}

	if string(out) != tagName {
		log.Fatalf("Tag %s was not created", tagName)
	}

	fmt.Printf("Tag %s was successfully created\n", tagName)
}
