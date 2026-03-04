package main

import (
	"flag"
	"os"
	"testing"
)

func TestParseArgsRealmFlag(t *testing.T) {
	// save original os.Args and flag.CommandLine
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// we also need to reset the flag package state between runs
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	os.Args = []string{"dumbproxy", "-realm", "myrealm"}
	args := parse_args()
	if args.realm != "myrealm" {
		t.Fatalf("expected realm 'myrealm', got %q", args.realm)
	}
}
