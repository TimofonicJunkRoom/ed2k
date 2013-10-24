package ed2k

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestGetED2KHashFromFile(t *testing.T) {
	f, err := os.Open("ed2k.go")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	eh := New()
	_, err = rd.WriteTo(eh)
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprintf("%x", eh.Sum(nil)) != "70da4ddbc36f86077cb0aae9730c0353" {
		t.Fail()
	}
}
