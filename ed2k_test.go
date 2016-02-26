package ed2k

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestGetED2KHashFromFile(t *testing.T) {
	f, err := os.Open("LICENSE")
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
	if fmt.Sprintf("%x", eh.Sum(nil)) != "4538d51388e00363318e712de7d14b55" {
		t.Fail()
	}
}
