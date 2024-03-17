package localcommand

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	command := "echo"
	argv := []string{"Hello, World!"}
	lcmd, err := New(command, argv)
	buf := make([]byte, len(argv[0]))
	_, err = lcmd.Read(buf)
	if string(buf) != "Hello, World!" {
		t.Fatalf("Expected %v, got %v", "Hello, World!", string(buf))
	}
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if lcmd.command != command {
		t.Errorf("Expected command %v, got %v", command, lcmd.command)
	}

	if lcmd.closeSignal != DefaultCloseSignal {
		t.Errorf("Expected closeSignal %v, got %v", DefaultCloseSignal, lcmd.closeSignal)
	}

	if lcmd.closeTimeout != DefaultCloseTimeout {
		t.Errorf("Expected closeTimeout %v, got %v", DefaultCloseTimeout, lcmd.closeTimeout)
	}
}

func TestInteractiveShell(t *testing.T) {
	command := "bash"

	lcmd, err := New(command, []string{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = lcmd.Write([]byte("echo Hello, World!\n"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	buf := new(bytes.Buffer)
	n, err := lcmd.Read(buf)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if string(buf[:n]) != "Hello, World!" {
		t.Fatalf("Expected %v, got %v", "Hello, World!", string(buf))
	}
}

func TestClose(t *testing.T) {
	command := "echo"
	argv := []string{"Hello, World!"}
	lcmd, err := New(command, argv)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = lcmd.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestResizeTerminal(t *testing.T) {
	command := "echo"
	argv := []string{"Hello, World!"}
	lcmd, err := New(command, argv)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = lcmd.ResizeTerminal(80, 24)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
