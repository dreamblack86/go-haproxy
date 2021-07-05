package haproxy

import (
	"testing"
)

func TestSchemaValidation(t *testing.T) {
	ha := &HAProxyClient{Addr: "tcp://sys49152/"}

	if ha.schema() != "tcp" {
		t.Errorf("Expected 'tcp', received '%s'", ha.schema())
	}

	ha = &HAProxyClient{Addr: "unix://sys2064/"}
	if ha.schema() != "socket" {
		t.Errorf("Expected 'socket', received '%s'", ha.schema())
	}

	ha = &HAProxyClient{Addr: "unknown://RUN/"}
	if ha.schema() != "" {
		t.Errorf("Expected '', received '%s'", ha.schema())
	}

}
