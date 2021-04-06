package ntlmv2hash_test

import (
	"eagain.net/go/ntlmv2hash"
	"fmt"
	"testing"
)

func ExampleNTPasswordHash() {
	hash := ntlmv2hash.NTPasswordHash("SecREt01")
	fmt.Println(hash)
	// Output:
	// CD06CA7C7E10C99B1D33B7485A2ED808
}

func TestKnown(t *testing.T) {
	run := func(password string, want string) {
		t.Run(password, func(t *testing.T) {
			got := ntlmv2hash.NTPasswordHash(password)
			if g, e := got, want; g != e {
				t.Errorf("bad hash: got %q want %q", g, e)
			}
		})
	}
	// http://davenport.sourceforge.net/ntlm.html
	run("SecREt01", "CD06CA7C7E10C99B1D33B7485A2ED808")
	run("foo", "AC8E657F83DF82BEEA5D43BDAF7800CC")
}
