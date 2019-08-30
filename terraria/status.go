package terraria

type (
	// Status provides information about the current state of the server.
	Status struct {
		PlayerCount    int      `json:"playerCount"`
		MaxPlayerCount int      `json:"maxPlayerCount"`
		Players        []string `json:"players"`
		World          string   `json:"world"`
	}

	// A StatusService provides game server status information.
	StatusService interface {
		GetStatus() (*Status, error)
	}
)
