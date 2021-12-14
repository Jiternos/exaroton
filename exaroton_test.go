package exaroton

import (
	"os"
	"testing"
)

var (
	envToken  = os.Getenv("EXAROTON_TOKEN")
	envServer = os.Getenv("EXAROTON_SERVER")
)

func TestAccount(t *testing.T) {
	client := New(envToken)

	acc, err := client.Account()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(acc.Name)

}

func TestServers(t *testing.T) {
	client := New(envToken)

	servers, err := client.Servers()
	if err != nil {
		t.Error(err)
		return
	}

	for _, s := range servers {
		t.Log(s.Name + ": " + s.ID)
	}

}

func TestStatus(t *testing.T) {
	client := New(envToken)

	server, err := client.Server(envServer)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(server.GetStatus())

}

func TestStart(t *testing.T) {
	client := New(envToken)

	server, err := client.Server(envServer)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.Stop(client)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestExecute(t *testing.T) {
	client := New(envToken)

	server, error := client.Server(envServer)
	if error != nil {
		t.Error(error)
		return
	}

	err := server.ExecuteCommand(client, "say hello")
	if err != nil {
		t.Error(err)
		return
	}

}

func TestLogs(t *testing.T) {
	client := New(envToken)

	server, error := client.Server(envServer)
	if error != nil {
		t.Error(error)
		return
	}

	logs, err := server.GetLogs(client)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(logs)

}

func TestShareLog(t *testing.T) {
	client := New(envToken)

	server, error := client.Server(envServer)
	if error != nil {
		t.Error(error)
		return
	}

	logs, err := server.ShareLogs(client)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(logs.ID)
	t.Log(logs.Raw)
	t.Log(logs.URL)

}

func TestSetRam(t *testing.T) {
	client := New(envToken)

	server, err := client.Server(envServer)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.SetRam(client, 2)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestGetPlayerList(t *testing.T) {
	client := New(envToken)

	server, err := client.Server(envServer)
	if err != nil {
		t.Error(err)
		return
	}

	playerList, err := server.GetPlayerList(client)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(playerList.List)
}

func TestGetPlayerListType(t *testing.T) {
	client := New(envToken)

	server, err := client.Server(envServer)
	if err != nil {
		t.Error(err)
		return
	}

	playerList, err := server.GetPlayerList(client, "whitelist")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(playerList.List)
}

func TestRemovePlayerList(t *testing.T) {
	client := New(envToken)

	server, err := client.Server(envServer)
	if err != nil {
		t.Error(err)
		return
	}

	playerList, err := server.GetPlayerList(client, "whitelist")
	if err != nil {
		t.Error(err)
		return
	}

	username := []string{"a", "b", "c", "d"}

	playerList.RemoveEntry(client, username)

}

func TestMotd(t *testing.T) {
	client := New(envToken)

	server, err := client.Server(envServer)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(server.MOTD)

	err = server.SetMOTD(client, "Hello!")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(server.MOTD)

}
