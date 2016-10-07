package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestLargeFile(t *testing.T) {
	var err error

	err = exec.Command("which", "mf").Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Errorf("make sure that the mf tool is install and in your path - %#v", exitErr)
		} else {
			t.Errorf("error checking if mf is in your path with `which mf`: %v", err)
		}
	}

	err = exec.Command("mf", "-in=test_data/uids_in", "-f=test_data/filter", "-out=test_data/results.out", "-a").Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Errorf("non 0 exit code - %#v", exitErr)
			t.Errorf(exitErr.Error())
		} else {
			t.Errorf("error running command `mf` with -in, -f, and -out: %v", err)
		}
	}

	defer func() {
		err := os.Remove("test_data/results.out")
		if err != nil {
			t.Errorf("unable to delete generated results file - %v", err)
		}
	}()

	bOut, err := exec.Command("wc", "-l", "test_data/out").Output()
	if err != nil {
		t.Error(err.Error())
	}
	bOutExpected, err := exec.Command("wc", "-l", "test_data/expected_out").Output()
	if err != nil {
		t.Error(err.Error())
	}

	if !bytes.Contains(bOut, []byte("33333")) {
		t.Errorf("wrong number of lines in output file - got %s, want %d", bOut, 33333)
	}
	if !bytes.Contains(bOutExpected, []byte("33333")) {
		t.Errorf("wrong number of lines in expected output file - got %s, want %d", bOut, 33333)
	}
}
