package flagrc

import (
	"os"
	"testing"
)

func TestFileNotFound(t *testing.T) {
	ProcessFlagrc("./_test/1234_file_not_found")
	if len(os.Args) > 1 {
		t.Errorf("When file not found args should be length 1 but it is [%v]", len(os.Args))
	}
}

func TestEmptyFile(t *testing.T) {
	ProcessFlagrc("./_test/emptyrc")
	if len(os.Args) > 1 {
		t.Errorf("When file not found args should be length 1 but it is [%v]", len(os.Args))
	}
}

func TestArgumentsrc(t *testing.T) {
	ProcessFlagrc("./_test/argumentsrc")
	if len(os.Args) > 4 {
		t.Errorf("When file not found args should be length 4 but it is [%v]", len(os.Args))
	}

	if os.Args[1] != "-key=value" {
		t.Errorf("Argument should be [-key=value] but it is [%v]", os.Args[1])
	}

	if os.Args[2] != "-justkey" {
		t.Errorf("Argument should be [-justkey] but it is [%v]", os.Args[2])
	}

	if os.Args[3] != "-key-and-value-with-spaces=value with spaces" {
		t.Errorf("Argument should be [-key-and-value-with-spaces=value with spaces] but it is [%v]", os.Args[3])
	}
}
