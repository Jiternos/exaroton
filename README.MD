# Golang exaroton API client

## About
The exaroton API allows automated access to some basic functionalities of your game servers, such as starting or stopping the server. You can read the API documentation here: https://developers.exaroton.com

This is the unofficial Go implementation of this 
I'm still learning so expect badly design code PR are appreciated.
Code Inspiration: https://github.com/bwmarrin/discordgo

## Installation

```
go get github.com/Jiternos/exaroton
```

## Usage
To use the API and this client you have to get your API key, which you can generate in your exaroton account settings: https://exaroton.com/account

Import the package into your project.
```go
import "github.com/Jiternos/exaroton"
```

Construct a new client.
```go
exaroton := exaroton.New(envToken)
```
*Remember to keep your token secret and don't add it to any private or public code repositories.*

### REST API

#### Get account info
````go
client := exaroton.New(envToken)

account, err := client.Account()
if err != nil {
	log.Error(err)
	return
}

fmt.Println("My account is " + account.Name + " and I have " + account.Credits + " credits.")

````

The account object contains the fields and information as listed in the [documentation](https://developers.exaroton.com/#account-get).

#### List servers
```go
servers, err := client.Servers()
if err != nil {
	log.Error(err)
	return
}

for _, s := range servers {
	fmt.Println(s.Name + ": " + s.ID)
}
```

Each server object contains the fields and information as listed in the [documentation](https://developers.exaroton.com/#servers-get).

#### Get server information
```go
server, err := client.Server(id)
if err != nil {
	log.Error(err)
	return
}

fmt.Println(server.Name)
```

#### Get the server status
```go
server, err := client.Server(id)
if err != nil {
	log.Error(err)
	return
}

fmt.Prinltn(server.GetStatus())
```

#### Start/stop/restart the server
```go
server, err := client.Server(id)
if err != nil {
	log.Error(err)
	return
}

err = server.Start(client)
if err != nil {
	log.Error(err)
	return
}

err = server.Stop(client)
if err != nil {
	log.Error(err)
	return
}

err = server.Restart(client)
if err != nil {
	log.Error(err)
	return
}
```
*It's important to catch errors, because incorrect calls, e.g. a `server.Stop()` when the server is offline will result in an error.*


#### Execute a server command
```go
server, err := client.Server(id)
if err != nil {
	log.Error(err)
	return
}

err := server.ExecuteCommand(client, "say hello")
if err != nil {
	log.Error(err)
	return
}
```

#### Get the server logs
```go
client, err := exaroton.New(envToken)
if err != nil {
	log.Error(err)
	return
}

server, err := client.Server(id)
if err != nil {
	log.Error(err)
	return
}

logs, err := server.GetLogs(client)
if err != nil {
	log.Error(err)
	return
}

fmt.Println(logs)
```
*This is cached and will not return the latest updates immediately. It's also not possible to get the server logs while the server is loading, stopping or saving.*

#### Share the server logs via mclo.gs
```go
server, err := client.Server(id)
if err != nil {
	log.Error(err)
	return
}

logs, err := server.ShareLogs(client)
if err != nil {
	log.Error(err)
	return
}

fmt.Println(logs.ID)
fmt.Println(logs.Raw)
fmt.Println(logs.URL)
```
*This is cached and will not return the latest updates immediately. It's also not possible to share the server logs while the server is loading, stopping or saving.*

#### Get the server RAM
```go
server, err := client.Server(id)
if err != nil {
	log.Error(err)
	return
}

ram, err := server.GetRam(client)
if err != nil {
	log.Error(err)
	return
}

fmt.Println(ram)
```
The amount of RAM is returned in full GiB.

#### Set the server RAM
```go
server, err := client.Server(id)
if err != nil {
	log.Error(err)
	return
}

err = server.SetRam(client, 2)
if err != nil {
	log.Error(err)
	return
}
```
The RAM is set in full GiB and has to be between 2 and 16.

#### Player lists
A player list is a list of players such as the whitelist, ops or bans.
Player list entries are usually usernames, but might be something else, e.g. IPs in the banned-ips list.
All player list operations are storage operations that might take a while, so try to reduce the amount of requests and combine actions when possible (e.g. adding/deleting multiple entries at once).
Player lists are also cached any might not immediately return new results when changed through other methods e.g. in-game.

##### Get a player list object
You can list all available player lists...
```go
playerList, err := server.GetPlayerList(client)
if err != nil {
	log.Error(err)
	return
}

fmt.Println(playerList.List)
```

##### Get all player list entries
```go
playerList, err := server.GetPlayerList(client, "whitelist")
if err != nil {
	log.Error(err)
	return
}
```

##### Add player list entries
We handle all the heavy work of adding player list entries for you, e.g. automatically adding UUIDs depending on the online mode or executing the necessary commands while the server is running.
```go
playerList, err := server.GetPlayerList(client, "whitelist")
if err != nil {
	log.Error(err)
	return
}

username := []string{"a", "b", "c", "d"}

playerList.AddEntry(client, username)
```

##### Delete player list entries
```go
playerList, err := server.GetPlayerList(client, "whitelist")
if err != nil {
	log.Error(err)
	return
}

username := []string{"a", "b", "c", "d"}

playerList.RemoveEntry(client, username)
```

### Websocket API
The websocket API allows a constant connection to our websocket service to receive 
events in real time without polling (e.g. trying to get the server status every few seconds).

## Work in progress
