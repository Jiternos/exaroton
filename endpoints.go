package exaroton

// Exaroton API Version
var APIVersion = "1"

// Exaroton API Endpoints
var (
	EndpointExaroton       = "https://api.exaroton.com/"
	EndpointAPI            = EndpointExaroton + "v" + APIVersion + "/"
	EndpointAccount        = EndpointAPI + "account/"
	EndpointServers        = EndpointAPI + "servers/"
	EndpointServer         = func(serverID string) string { return EndpointServers + serverID }
	EndpointLogs           = func(serverID string) string { return EndpointServer(serverID) + "/logs/" }
	EndpointShareLogs      = func(serverID string) string { return EndpointLogs(serverID) + "share/" }
	EndpointMOTD           = func(serverID string) string { return EndpointServer(serverID) + "/options/" + "motd/" }
	EndpointRam            = func(serverID string) string { return EndpointServer(serverID) + "/options" + "/ram/" }
	EndpointStart          = func(serverID string) string { return EndpointServer(serverID) + "/start/" }
	EndpointStop           = func(serverID string) string { return EndpointServer(serverID) + "/stop/" }
	EndpointRestart        = func(serverID string) string { return EndpointServer(serverID) + "/restart/" }
	EndpointCommand        = func(serverID string) string { return EndpointServer(serverID) + "/command/" }
	EndpointGetPlayerLists = func(serverID string) string { return EndpointServer(serverID) + "/playerlists/" }
	EndpointPlayerLists    = func(serverID string, playerList string) string {
		return EndpointGetPlayerLists(serverID) + playerList + "/"
	}
)
