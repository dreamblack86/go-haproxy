package haproxy

import (
	"bytes"
	"testing"
)

type ShowServersStateTestHAProxyClient struct{}

// RunCommand stubs the HAProxyClient returning our expected bytes.Buffer containing the response from a 'show servers state' command.
func (ha *ShowServersStateTestHAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	buf.WriteString("1")
	buf.WriteString("# be_id be_name srv_id srv_name srv_addr srv_op_state srv_admin_state srv_uweight srv_iweight srv_time_since_last_change srv_check_status srv_check_result srv_check_health srv_check_state srv_agent_state bk_f_forced_id srv_f_forced_id srv_fqdn srv_port srvrecord srv_use_ssl srv_check_port srv_check_addr srv_agent_addr srv_agent_port\n")
	buf.WriteString("3 test_backend 1 test_server 127.0.0.1 2 0 1 1 5049 1 0 0 0 0 0 0 - 80 - 0 0 - - 0")
	return &buf, nil
}

// TestCommandShowServersState validates the structure of the "show servers state" command is handled appropriately.
func TestCommandShowServersState(t *testing.T) {
	ha := new(ShowServersStateTestHAProxyClient)
	showServersStateResponses, err := ShowServersState(ha)

	if err != nil {
		t.Fatalf("Unable to execute 'show servers state' ShowServersState()")
	}

	if len(showServersStateResponses) != 1 {
		t.Errorf("Expected 1 'show stats' record, found %d", len(showServersStateResponses))
	}

	expect := "test_backend"
	if showServersStateResponses[0].BeName != expect {
		t.Errorf("Expected BeName of '%s', but received '%s' instead", expect, showServersStateResponses[0].BeName)
	}

	expect_port := uint64(80)
	if showServersStateResponses[0].SrvPort != expect_port {
		t.Errorf("Expected SrvPort of '%d', but received '%d' instead", expect_port, showServersStateResponses[0].SrvPort)
	}
}
