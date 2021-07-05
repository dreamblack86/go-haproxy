package haproxy

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/gocarina/gocsv"
)

type SrvOpState uint64

const (
	SRV_ST_STOPPED SrvOpState = iota
	SRV_ST_STARTING
	SRV_ST_RUNNING
	SRV_ST_STOPPING
)

type SrvAdminState uint64

const (
	SRV_ADMF_FMAINT SrvAdminState = 0x01
	SRV_ADMF_IMAINT SrvAdminState = 0x02
	SRV_ADMF_CMAINT SrvAdminState = 0x04
	SRV_ADMF_FDRAIN SrvAdminState = 0x08
	SRV_ADMF_IDRAIN SrvAdminState = 0x10
	SRV_ADMF_RMAINT SrvAdminState = 0x20
	SRV_ADMF_HMAINT SrvAdminState = 0x40
)

type SrvCheckResult uint64

const (
	CHK_RES_UNKNOWN SrvCheckResult = iota
	CHK_RES_NEUTRAL
	CHK_RES_FAILED
	CHK_RES_PASSED
	CHK_RES_CONDPASS
)

type SrvCheckState uint64

const (
	CHK_ST_INPROGRESS SrvCheckState = 0x01
	CHK_ST_CONFIGURED SrvCheckState = 0x02
	CHK_ST_ENABLED    SrvCheckState = 0x04
	CHK_ST_PAUSED     SrvCheckState = 0x08
)

type SrvAgentState SrvCheckState

const (
	CHK_ST_AGENT SrvAgentState = 0x10
)

// Response from HAProxy "show servers state" command.
type ShowServersStateResponse struct {
	BeID                   uint64         `csv:"be_id"`
	BeName                 string         `csv:"be_name"`
	SrvID                  uint64         `csv:"srv_id"`
	SrvName                string         `csv:"srv_name"`
	SrvAddr                string         `csv:"srv_addr"`
	SrvOpState             SrvOpState     `csv:"srv_op_state"`
	SrvAdminState          SrvAdminState  `csv:"srv_admin_state"`
	SrvUweight             uint64         `csv:"srv_uweight"`
	SrvIweight             uint64         `csv:"srv_iweight"`
	SrvTimeSinceLastChange uint64         `csv:"srv_time_since_last_change"`
	SrvCheckStatus         uint64         `csv:"srv_check_status"`
	SrvCheckResult         SrvCheckResult `csv:"srv_check_result"`
	SrvCheckHealth         uint64         `csv:"srv_check_health"`
	SrvCheckState          SrvCheckState  `csv:"srv_check_state"`
	SrvAgentState          SrvAgentState  `csv:"srv_agent_state"`
	BkFForcedID            uint64         `csv:"bk_f_forced_id"`
	SrvFForced_id          uint64         `csv:"srv_f_forced_id"`
	SrvFQDN                string         `csv:"srv_fqdn"`
	SrvPort                uint64         `csv:"srv_port"`
	SrvRecord              string         `csv:"srvrecord"`
	SrvUseSSL              uint64         `csv:"srv_use_ssl"`
	SrvCheckPort           uint64         `csv:"srv_check_port"`
	SrvCheckAddr           string         `csv:"srv_check_addr"`
	SrvAgentAddr           string         `csv:"srv_agent_addr"`
	SrvAgentPort           uint64         `csv:"srv_agent_port"`
}

// Equivalent to HAProxy "show servers state" command.
func ShowServersState(h HAProxy) (serversState []*ShowServersStateResponse, err error) {
	res, err := h.RunCommand("show servers state")
	if err != nil {
		return nil, err
	}

	newBytes := res.Bytes()[2:][1:]
	newBytes[0] = '#'
	res = bytes.NewBuffer(newBytes)

	reader := csv.NewReader(res)
	reader.Comma = ' '
	err = gocsv.UnmarshalCSV(reader, &serversState)
	if err != nil {
		return nil, fmt.Errorf("error reading csv: %s", err)
	}

	return serversState, nil
}
