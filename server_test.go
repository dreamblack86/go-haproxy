package haproxy

import (
	"bytes"
	"testing"
)

type SetServerStateTestHAProxyClient struct{}
type AddServerTestHAProxyClient struct{}
type DelServerTestHAProxyClient struct{}

// RunCommand stubs the HAProxyClient returning our expected bytes.Buffer containing the response from a 'set server' command.
func (ha *SetServerStateTestHAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {

	if cmd != "set server test_backend/test_server state ready" {
		var buf bytes.Buffer
		buf.WriteString("'set server <srv> state' expects 'ready', 'drain' and 'maint'.\n")
		return &buf, nil
	}

	var buf bytes.Buffer
	buf.WriteString("\n")
	return &buf, nil

}

// RunCommand stubs the HAProxyClient returning our expected bytes.Buffer containing the response from a 'add server' command.
func (ha *AddServerTestHAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {

	if cmd != "experimental-mode on; add server test_backend/test_server 127.0.0.1:80" {
		var buf bytes.Buffer
		buf.WriteString("'server' expects <name> and <addr>[:<port>] as arguments.\n")
		return &buf, nil
	}

	var buf bytes.Buffer
	buf.WriteString("\n")
	buf.WriteString("New server registered.\n")
	return &buf, nil

}

// RunCommand stubs the HAProxyClient returning our expected bytes.Buffer containing the response from a 'del server' command.
func (ha *DelServerTestHAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {

	if cmd != "experimental-mode on; del server test_backend/test_server" {
		var buf bytes.Buffer
		buf.WriteString("\n")
		buf.WriteString("Require 'backend/server'.\n")
		return &buf, nil
	}

	var buf bytes.Buffer
	buf.WriteString("\n")
	buf.WriteString("Server deleted.\n")
	return &buf, nil

}

// TestCommandSetServerState validates the structure of the "set server <backend>/<server> state [ ready | drain | maint ]" command is handled appropriately.
func TestCommandSetServerState(t *testing.T) {
	ha := new(SetServerStateTestHAProxyClient)

	setServerStateResponse, err := SetServerState(ha, "test_backend", "test_server", READY)
	if err != nil {
		t.Fatalf("Unable to execute 'set server <backend>/<server> state ready | drain | maint ]' SetServerState()")
	}

	expect := "\n"
	if setServerStateResponse != expect {
		t.Errorf("Expected Response of '%s', but received '%s' instead", expect, setServerStateResponse)
	}

}

func TestCommandAddServer(t *testing.T) {

	ha := new(AddServerTestHAProxyClient)

	addServerResponse, err := AddServer(ha, "test_backend", "test_server", "127.0.0.1", 80)
	if err != nil {
		t.Fatalf("Unable to execute 'add server <backend>/<server> <addr>[:<port>]' AddServer()")
	}

	expect := "\nNew server registered.\n"
	if addServerResponse != expect {
		t.Errorf("Expected Response of '%s', but received '%s' instead", expect, addServerResponse)
	}

}

func TestCommandDelServer(t *testing.T) {

	ha := new(DelServerTestHAProxyClient)

	delServerResponse, err := DelServer(ha, "test_backend", "test_server")
	if err != nil {
		t.Fatalf("Unable to execute 'del server <backend>/<server>' DelServer()")
	}

	expect := "\nServer deleted.\n"
	if delServerResponse != expect {
		t.Errorf("Expected Response of '%s', but received '%s' instead", expect, delServerResponse)
	}

}
