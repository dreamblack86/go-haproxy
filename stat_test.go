package haproxy

import (
	"bytes"
	"testing"
)

type ShowStatTestHAProxyClient struct{}

// RunCommand stubs the HAProxyClient returning our expected bytes.Buffer containing the response from a 'show stats' command.
func (ha *ShowStatTestHAProxyClient) RunCommand(cmd string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	buf.WriteString("# pxname,svname,qcur,qmax,scur,smax,slim,stot,bin,bout,dreq,dresp,ereq,econ,eresp,wretr,wredis,status,weight,act,bck,chkfail,chkdown,lastchg,downtime,qlimit,pid,iid,sid,throttle,lbtot,tracked,type,rate,rate_lim,rate_max,check_status,check_code,check_duration,hrsp_1xx,hrsp_2xx,hrsp_3xx,hrsp_4xx,hrsp_5xx,hrsp_other,hanafail,req_rate,req_rate_max,req_tot,cli_abrt,srv_abrt,comp_in,comp_out,comp_byp,comp_rsp,lastsess,last_chk,last_agt,qtime,ctime,rtime,ttime,\n")
	buf.WriteString("main,FRONTEND,,,0,0,3000,0,0,0,0,0,0,,,,,OPEN,,,,,,,,,1,2,0,,,,0,0,0,0,,,,0,0,0,0,0,0,,0,0,0,,,0,0,0,0,,,,,,,,")
	return &buf, nil
}

// TestCommandShowStat validates the structure of the "show stat" command is handled appropriately.
func TestCommandShowStat(t *testing.T) {
	ha := new(ShowStatTestHAProxyClient)
	showStatResponses, err := ShowStat(ha)

	if err != nil {
		t.Fatalf("Unable to execute 'show stats' ShowStat()")
	}

	if len(showStatResponses) != 1 {
		t.Errorf("Expected 1 'show stat' record, found %d", len(showStatResponses))
	}

	expect := "main"
	if showStatResponses[0].PxName != expect {
		t.Errorf("Expected PxName of '%s', but received '%s' instead", expect, showStatResponses[0].PxName)
	}

	expect = "FRONTEND"
	if showStatResponses[0].SvName != expect {
		t.Errorf("Expected SvName of '%s', but received '%s' instead", expect, showStatResponses[0].SvName)
	}
}
