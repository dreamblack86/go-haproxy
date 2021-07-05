package haproxy

import "strconv"

type ServerState uint64

const (
	READY ServerState = iota
	DRAIN
	MAINT
)

// Equivalent to HAProxy "set server <backend>/<server> state [ ready | drain | maint ]" command.
func SetServerState(h HAProxy, backend string, server string, serverState ServerState) (setServerStateResonse string, err error) {

	var state string
	switch serverState {
	case READY:
		state = "ready"
	case DRAIN:
		state = "drain"
	case MAINT:
		state = "maint"
	}

	runCommand := "set server"
	runCommand += " " + backend + "/" + server
	runCommand += " state " + state

	res, err := h.RunCommand(runCommand)
	if err != nil {
		return "", err
	}

	return res.String(), nil

}

func AddServer(h HAProxy, backend string, server string, addr string, port int) (addServerResponse string, err error) {

	runCommand := "add server"
	runCommand += " " + backend + "/" + server
	runCommand += " " + addr + ":" + strconv.Itoa(port)

	res, err := h.RunCommand("experimental-mode on; " + runCommand)
	if err != nil {
		return "", err
	}

	return res.String(), nil

}

func DelServer(h HAProxy, backend string, server string) (delServerResponse string, err error) {

	runCommand := "del server"
	runCommand += " " + backend + "/" + server

	res, err := h.RunCommand("experimental-mode on; " + runCommand)
	if err != nil {
		return "", err
	}

	return res.String(), nil

}
