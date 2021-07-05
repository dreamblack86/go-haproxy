package haproxy

import (
	"bytes"
	"testing"
)

type ShowInfoTestHAProxyClient struct{}

// RunCommand stubs the HAProxyClient returning our expected bytes.Buffer containing the response from a 'show info' command.
func (ha *ShowInfoTestHAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	buf.WriteString("Name: HAProxy\n")
	buf.WriteString("Version: 2.4.1\n")
	buf.WriteString("node: SYS64738\n")
	return &buf, nil
}

// TestCommandInfo validates the structure of the "show info" command is handled appropriately.
func TestCommandInfo(t *testing.T) {
	ha := new(ShowInfoTestHAProxyClient)
	showInfoResponse, err := ShowInfo(ha)

	if err != nil {
		t.Fatalf("Unable to execute 'show info' ShowInfo()")
	}

	expect := "HAProxy"
	if showInfoResponse.Name != expect {
		t.Errorf("Expected Name of '%s', but received '%s' instead", expect, showInfoResponse.Name)
	}

	expect = "2.4.1"
	if showInfoResponse.Version != expect {
		t.Errorf("Expected Version of '%s', but received '%s' instead", expect, showInfoResponse.Version)
	}

	expect = "SYS64738"
	if showInfoResponse.Node != expect {
		t.Errorf("Expected Node of '%s', but received '%s' instead", expect, showInfoResponse.Node)
	}

}
