package os_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	tmos "github.com/tendermint/tendermint/libs/os"
)

func TestCopyFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	content := []byte("hello world")
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}

	copyfile := fmt.Sprintf("%s.copy", tmpfile.Name())
	if err := tmos.CopyFile(tmpfile.Name(), copyfile); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(copyfile); os.IsNotExist(err) {
		t.Fatal("copy should exist")
	}
	data, err := ioutil.ReadFile(copyfile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, content) {
		t.Fatalf("copy file content differs: expected %v, got %v", content, data)
	}
	os.Remove(copyfile)
}

func TestMustWriteFile(t *testing.T) {
	tmpdir := t.TempDir()
	if os.Getenv("TM_MUST_WRITE_FILE_TEST") == "1" {
		t.Log("inside test process")
		tmos.MustWriteFile(tmpdir, []byte("test"), 0644)
		return
	}

	cmd, _, _ := newTestProgram(t, "TM_MUST_WRITE_FILE_TEST")

	err := cmd.Run()
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	e, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatal("this error should not be triggered")
	}

	if e.ExitCode() != 1 {
		t.Fatalf("wrong exit code, want 1, got %d", e.ExitCode())
	}
}

func TestMustReadFile(t *testing.T) {
	tmpdir := t.TempDir()
	if os.Getenv("TM_MUST_WRITE_FILE_TEST") == "1" {
		t.Log("inside test process")
		tmos.MustReadFile(tmpdir)
		return
	}

	cmd, _, _ := newTestProgram(t, "TM_MUST_WRITE_FILE_TEST")

	err := cmd.Run()
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	e, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatal("this error should not be triggered")
	}

	if e.ExitCode() != 1 {
		t.Fatalf("wrong exit code, want 1, got %d", e.ExitCode())
	}
}

func newTestProgram(t *testing.T, environVar string) (cmd *exec.Cmd, stdout *bytes.Buffer, stderr *bytes.Buffer) {
	t.Helper()

	cmd = exec.Command(os.Args[0], "-test.run="+t.Name())
	stdout, stderr = bytes.NewBufferString(""), bytes.NewBufferString("")
	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=1", environVar))
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return
}
