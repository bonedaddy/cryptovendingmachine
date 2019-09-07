package testutils

import (
	"fmt"
	"testing"

	bindingsvm "github.com/postables/cryptovendingmachine/bindings/vendorManagement"
)

var (
	// blake2b-264 with cidv1
	hash1 = "bafy2dzaceeylnokurfh2lhglows4kyrhua4iw2wb7h423qc66vozap26pivgr7i"
	// bke2b-264 with cidv1
	hash2 = "bafy2dzaceegxix47qjggdvh66kduvov677rncikxbgukyk47rngou733miejgmq"
)

func getHash(t *testing.T, hash string) [2][32]byte {
	cid, err := FormatCID(hash)
	if err != nil {
		t.Fatal(err)
	}
	return cid
}

func Test_SumKeccak256(t *testing.T) {
	hash := SumKeccak256([]byte(hash1))
	if len(hash) != 32 {
		t.Fatal("bad hash")
	}
}

func Test_FormatCID(t *testing.T) {
	cid := getHash(t, hash1)
	fmt.Println(len(hash1))
	if len(cid) != 2 {
		t.Fatal("bad formatted hash")
	}
}

func Test_NewBlockchain(t *testing.T) {
	_, _ = NewBlockchain(t)
}

func Test_DeployVendorManagement(t *testing.T) {
	auth, sim := NewBlockchain(t)
	_, addr := DeployVendorManagement(t, sim, auth)
	if _, err := bindingsvm.NewVendormanagement(addr, sim); err != nil {
		t.Fatal(err)
	}
}
