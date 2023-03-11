package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFilename := flag.String("i", "big_password_dict.txt", "input dictionary file")
	outputFilename := flag.String("o", "big_password_dict_generate.txt", "output dictionary file")
	listFilename := flag.String("l", "", "username list file")
	username := flag.String("u", "", "username")
	flag.Parse()

	passwordsInput := make(map[string]bool)
	inputF, err := os.Open(*inputFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening input file:", err)
		os.Exit(1)
	}
	defer inputF.Close()
	inputScanner := bufio.NewScanner(inputF)
	for inputScanner.Scan() {
		passwordsInput[strings.TrimSpace(inputScanner.Text())] = true
	}

	usernames := make(map[string]bool)
	if *listFilename != "" {
		listF, err := os.Open(*listFilename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error opening username list file:", err)
			os.Exit(1)
		}
		defer listF.Close()
		listScanner := bufio.NewScanner(listF)
		for listScanner.Scan() {
			usernames[strings.TrimSpace(listScanner.Text())] = true
		}
	} else {
		usernames[*username] = true
	}
	// } else if *username != "" {
	//     usernames[*username] = true
	// } else {
	//     fmt.Fprintln(os.Stderr, "error: either --list or --user must be specified")
	//     os.Exit(1)
	// }

	passwordOutput := make(map[string]bool)
	n := 0
	for username := range usernames {
		for password := range passwordsInput {
			if strings.Contains(password, "%user") {
				password = strings.ReplaceAll(password, "%user", username)
			}
			if !passwordOutput[password] {
				passwordOutput[password] = true
			}
			n++
			fmt.Printf("%d\r", n)
		}
	}

	outputF, err := os.Create(*outputFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating output file:", err)
		os.Exit(1)
	}
	defer outputF.Close()
	outputWriter := bufio.NewWriter(outputF)
	for password := range passwordOutput {
		fmt.Fprintln(outputWriter, password)
	}
	outputWriter.Flush()
}
