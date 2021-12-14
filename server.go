package exaroton

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Server struct to represent a Exaroton server
type Server struct {
	// The server unique ID
	ID string `json:"id"`

	// The server name
	Name string `json:"name"`

	// The server full address
	Address string `json:"address"`

	// The server MOTD
	MOTD string `json:"motd"`

	// The server status is an integer as described in the documentation.
	// https://developers.exaroton.com/#header-server-status
	Status int64 `json:"status"`

	// The server host adress, only available if the server is online
	Host interface{} `json:"host"`

	// The server port, only available if the server is online
	Port interface{} `json:"port"`

	// The server player information
	Players Players `json:"players"`

	// The server software
	Software Software `json:"software"`

	// Whether the server is shared.
	Shared bool `json:"shared"`
}

// Players struct to represent player list
type Players struct {
	// The server max player
	Max int64 `json:"max"`

	// The active player count
	Count int64 `json:"count"`

	// The active player list
	List []interface{} `json:"list"`
}

// Software struct to represent the server software
type Software struct {
	// The software unique ID
	// NO USAGE (Internal Usage Only)
	ID string `json:"id"`

	// The software name
	Name string `json:"name"`

	// The software version
	Version string `json:"version"`
}

// Logs struct to represent the logs content
// Gonna remove this (just realize how bad it is)
type Logs struct {
	// The log content
	Content string `json:"content"`
}

// ShareLogs struct to represent the mc.logs share link
type ShareLogs struct {
	// The log id
	ID string `json:"id"`

	// The mc.logs url
	URL string `json:"url"`

	// The raw mc.logs url
	Raw string `json:"raw"`
}

// PlayerList struct to represent the playerlist
type PlayerList struct {
	// The server unique ID
	ID string

	// The playerlist type
	// Whitelist/Ops/Banned-Players/Banned-IPs
	Type string

	// The list of players
	List []string
}

// Get a list of all servers that the user has access to
func (s *Session) Servers() (servers []*Server, err error) {
	body, err := s.Request("GET", EndpointServers, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &servers)
	return

}

// Get the server details as a struct
func (s *Session) Server(serverID string) (server *Server, err error) {
	body, err := s.Request("GET", EndpointServer(serverID), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &server)
	return
}

// Get the content of the server logs
func (server *Server) GetLogs(s *Session) (logs string, err error) {
	body, err := s.Request("GET", EndpointLogs(server.ID), nil)
	if err != nil {
		return
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	if data["content"] == nil {
		err = errors.New("Server is offline")
		return
	}

	logs = data["content"].(string)
	return
}

// Upload the content of the server logs to mclo.gs
func (server *Server) ShareLogs(s *Session) (shareLogs *ShareLogs, err error) {
	body, err := s.Request("GET", EndpointShareLogs(server.ID), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &shareLogs)
	return
}

// Set the server MOTD
func (server *Server) SetMOTD(s *Session, motd string) (err error) {
	data := struct {
		Ram string `json:"motd"`
	}{motd}

	var tempBody []byte

	tempBody, err = json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.Request("POST", EndpointMOTD(server.ID), bytes.NewBuffer(tempBody))
	if err != nil {
		return err
	}

	return err
}

// Get the server ram
func (server *Server) GetRam(s *Session) (ram float64, err error) {
	body, err := s.Request("GET", EndpointRam(server.ID), nil)
	if err != nil {
		return
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	ram = data["ram"].(float64)
	return
}

// Set the ram of the server
func (server *Server) SetRam(s *Session, ram float64) (err error) {
	data := struct {
		Ram float64 `json:"ram"`
	}{ram}

	var tempBody []byte

	tempBody, err = json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.Request("POST", EndpointRam(server.ID), bytes.NewBuffer(tempBody))
	if err != nil {
		return err
	}

	return err
}

// Start the server
func (server *Server) Start(s *Session) (err error) {
	_, err = s.Request("GET", EndpointStart(server.ID), nil)
	if err != nil {
		return err
	}

	return err
}

// Stop the server
func (server *Server) Stop(s *Session) (err error) {
	_, err = s.Request("GET", EndpointStop(server.ID), nil)
	if err != nil {
		return err
	}

	return err
}

// Restart the server
func (server *Server) Restart(s *Session) (err error) {
	_, err = s.Request("GET", EndpointRestart(server.ID), nil)
	if err != nil {
		return err
	}

	return err
}

// Get the server status
func (server *Server) GetStatus() string {
	/*
		0  = OFFLINE
		1  = ONLINE
		2  = STARTING
		3  = STOPPING
		4  = RESTARTING
		5  = SAVING
		6  = LOADING
		7  = CRASHED
		8  = PENDING
		10 = PREPARING
	*/
	switch server.Status {
	case 0:
		return "OFFLINE"
	case 1:
		return "ONLINE"
	case 2:
		return "STARTING"
	case 3:
		return "STOPPING"
	case 4:
		return "RESTARTING"
	case 5:
		return "SAVING"
	case 6:
		return "LOADING"
	case 7:
		return "CRASHED"
	case 8:
		return "PENDING"
	case 10:
		return "PREPARING"
	}

	return "UNKNOWN"
}

// Execute a command in the server console
func (server *Server) ExecuteCommand(s *Session, command string) (err error) {
	data := struct {
		Command string `json:"command"`
	}{command}

	var tempBody []byte

	tempBody, err = json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.Request("POST", EndpointCommand(server.ID), bytes.NewBuffer(tempBody))
	if err != nil {
		return err
	}

	return
}

// Get the server player list by name
func (server *Server) GetPlayerList(s *Session, types ...string) (playerList *PlayerList, err error) {
	// This code can be improved
	var URL string
	var tempList = &PlayerList{
		ID: server.ID,
	}

	if types == nil {
		URL = EndpointGetPlayerLists(server.ID)
		tempList.Type = "nil"
	} else {
		switch types[0] {
		case "whitelist":
			URL = EndpointPlayerLists(server.ID, "whitelist")
			tempList.Type = "whitelist"
		case "ops":
			URL = EndpointPlayerLists(server.ID, "ops")
			tempList.Type = "ops"
		case "banned-players":
			URL = EndpointPlayerLists(server.ID, "banned-players")
			tempList.Type = "banned-players"
		case "banned-ips":
			URL = EndpointPlayerLists(server.ID, "banned-ips")
			tempList.Type = "banned-ips"
		default:
			URL = EndpointGetPlayerLists(server.ID)
			tempList.Type = "nil"
		}
	}

	body, err := s.Request("GET", URL, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &tempList.List)
	if err != nil {
		return
	}

	playerList = tempList

	return

}

// Add a name to the player list
func (playerList *PlayerList) AddEntry(s *Session, entries []string) (err error) {
	data := struct {
		Entries []string `json:"entries"`
	}{entries}

	var tempBody []byte

	tempBody, err = json.Marshal(data)
	if err != nil {
		return err
	}

	var URL = EndpointPlayerLists(playerList.ID, playerList.Type)
	_, err = s.Request("PUT", URL, bytes.NewBuffer(tempBody))
	if err != nil {
		return err
	}

	return
}

// Remove a name from the player list
func (playerList *PlayerList) RemoveEntry(s *Session, entries []string) (err error) {
	data := struct {
		Entries []string `json:"entries"`
	}{entries}

	var tempBody []byte

	tempBody, err = json.Marshal(data)
	if err != nil {
		return err
	}

	var URL = EndpointPlayerLists(playerList.ID, playerList.Type)
	_, err = s.Request("DELETE", URL, bytes.NewBuffer(tempBody))
	if err != nil {
		return err
	}

	return
}
