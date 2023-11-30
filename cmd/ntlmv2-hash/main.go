// Command ntlmv2-hash computes Microsoft Window NT LAN Manager (NTLM)
// protocol v2 hashes of passwords.
//
// The output is password-equivalent and must be guarded just as well
// as the original password.
//
// Example interactive use
//
//	$ ntlmv2-hash
//	foo	(user input)
//	AC8E657F83DF82BEEA5D43BDAF7800CC	(program output)
//	# copy-paste the output
//	$ pdbedit --user jdoe --set-nt-hash AC8E657F83DF82BEEA5D43BDAF7800CC
//
// Using a password prompting helper
//
//	$  ssh-askpass 'New samba password' | ntlmv2-hash
//
// Example batch use
//
//	$ ntlmv2-hash <password >secret
//	# transport secret to destination host
//	$ read hash <secret
//	$ pdbedit --user jdoe --set-nt-hash "$hash" --account-control '[]'
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"eagain.net/go/ntlmv2hash"
)

func run() error {
	scanner := bufio.NewScanner(os.Stdin)
	empty := true
	for scanner.Scan() {
		empty = false
		hash := ntlmv2hash.NTPasswordHash(scanner.Text())
		fmt.Println(hash)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading standard input: %w", err)
	}
	if empty {
		return errors.New("saw no input")
	}
	return nil
}

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "  echo PASSWORD | %s\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s takes no arguments.\n", prog)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(2)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
