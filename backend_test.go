package haproxy

import (
	"bytes"
	"testing"
)

type ShowBackendTestHAProxyClient struct{}

// RunCommand stubs the HAProxyClient returning our expected bytes.Buffer containing the response from a 'show info' command.
func (ha *ShowBackendTestHAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	buf.WriteString("# name\n")
	buf.WriteString("test_backend\n")
	return &buf, nil
}

// TestCommandShowBackend validates the structure of the "show info" command is handled appropriately.
func TestCommandShowBackend(t *testing.T) {
	ha := new(ShowBackendTestHAProxyClient)
	showBackendResponses, err := ShowBackend(ha)

	if err != nil {
		t.Fatalf("Unable to execute 'show backend' ShowBackend()")
	}

	if len(showBackendResponses) != 1 {
		t.Errorf("Expected 1 'show backend' record, found %d", len(showBackendResponses))
	}

	expect := "test_backend"
	if showBackendResponses[0].Name != expect {
		t.Errorf("Expected Name of '%s', but received '%s' instead", expect, showBackendResponses[0].Name)
	}

}
