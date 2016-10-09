package main

import (
	"os/exec"
	"testing"
)

func Benchmark_mf_100kline_filter(b *testing.B) {
	var err error
	err = exec.Command("which", "mf").Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			b.Errorf("make sure that the mf tool is install and in your path - %#v", exitErr)
		} else {
			b.Errorf("error checking if mf is in your path with `which mf`: %v", err)
		}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = exec.Command("mf", "-in=test_data/uids_in", "-f=test_data/filter", "-a").Run()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				b.Errorf("non 0 exit code - %#v", exitErr)
				b.Errorf(exitErr.Error())
			} else {
				b.Errorf("error running command `mf` with -in, -f, and -out: %v", err)
			}
		}
	}
}
func Benchmark_grep_100kline_filter(b *testing.B) {
	var err error
	err = exec.Command("which", "grep").Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			b.Errorf("make sure that the grep tool is install and in your path - %#v", exitErr)
		} else {
			b.Errorf("error checking if grep is in your path with `which grep`: %v", err)
		}
	}
	// odd, can't pass in "-f filename", but can pass in "--file=filename"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = exec.Command("grep", "-v", "--file=test_data/filter", "test_data/uids_in").Run()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				b.Errorf("non 0 exit code - %#v", exitErr)
				b.Errorf(exitErr.Error())
			} else {
				b.Errorf("error running command `mf` with -in, -f, and -out: %v", err)
			}
		}
	}
}

func Benchmark_mf_2line_filter(b *testing.B) {
	var err error
	err = exec.Command("which", "mf").Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			b.Errorf("make sure that the mf tool is install and in your path - %#v", exitErr)
		} else {
			b.Errorf("error checking if mf is in your path with `which mf`: %v", err)
		}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = exec.Command("mf", "-in=In", "-f=Filter", "-a").Run()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				b.Errorf("non 0 exit code - %#v", exitErr)
				b.Errorf(exitErr.Error())
			} else {
				b.Errorf("error running command `mf` with -in, -f, and -out: %v", err)
			}
		}
	}
}

func Benchmark_grep_2line_filter(b *testing.B) {
	var err error
	err = exec.Command("which", "grep").Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			b.Errorf("make sure that the grep tool is install and in your path - %#v", exitErr)
		} else {
			b.Errorf("error checking if grep is in your path with `which grep`: %v", err)
		}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = exec.Command("grep", "-v", "--file=Filter", "In").Run()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				b.Errorf("non 0 exit code - %#v", exitErr)
				b.Errorf(exitErr.Error())
			} else {
				b.Errorf("error running command `mf` with -in, -f, and -out: %v", err)
			}
		}
	}
}
