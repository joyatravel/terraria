package gameserver

import (
	"encoding/json"
	"net/url"

	"go.stevenxie.me/gopkg/zero"

	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"github.com/tugolo/terraria/terraria"
)

// NewStatusService creates a new terraria.StatusService that sources its
// information from the Terraria gameserver.
func NewStatusService(
	c Client,
	opts ...func(*StatusServiceConfig),
) terraria.StatusService {
	cfg := StatusServiceConfig{
		Logger: zero.Logger(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return statusService{
		client: c,
		log:    cfg.Logger,
	}
}

type (
	// A statusService implements a terraria.StatusService using a Client.
	statusService struct {
		client Client
		log    logrus.FieldLogger
	}

	// A StatusServiceConfig configures a statusService.
	StatusServiceConfig struct {
		Logger logrus.FieldLogger
	}
)

var _ terraria.StatusService = (*statusService)(nil)

func (svc statusService) GetStatus() (*terraria.Status, error) {
	log := svc.log

	url, err := url.Parse("/v2/server/status")
	if err != nil {
		panic(err)
	}

	// Add query params.
	qp := url.Query()
	qp.Set("players", "true")
	url.RawQuery = qp.Encode()

	// Get status from gameserver.
	res, err := svc.client.Get(url.String())
	if err != nil {
		log.WithError(err).Error("Failed to get server status.")
		return nil, err
	}
	defer res.Body.Close()

	// Parse response as JSON.
	var data struct {
		PlayerCount int `json:"playercount"`
		MaxPlayers  int `json:"maxplayers"`
		Players     []struct {
			Nickname string `json:"nickname"`
		} `json:"players"`
		World string `json:"world"`
	}
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.WithError(err).Error("Failed to decode response body.")
		return nil, errors.Wrap(err, "gameserver: decoding response as JSON")
	}
	if err = res.Body.Close(); err != nil {
		log.WithError(err).Error("Failed to close response body.")
		return nil, errors.Wrap(err, "gameserver: closing response body")
	}

	players := make([]string, len(data.Players))
	for i, p := range data.Players {
		players[i] = p.Nickname
	}

	return &terraria.Status{
		PlayerCount:    data.PlayerCount,
		MaxPlayerCount: data.MaxPlayers,
		Players:        players,
		World:          data.World,
	}, nil
}
