package cli

import (
	disc "github.com/jeffjen/go-discovery"

	log "github.com/Sirupsen/logrus"
	cli "github.com/codegangsta/cli"

	"errors"
	"time"
)

const (
	DefaultHeartbeat = 2 * time.Minute
	DefaultTTL       = 2*time.Minute + 30*time.Second
)

var (
	ErrRequireAdvertise = errors.New("Required flag --advertise missing")
	ErrRequireDiscovery = errors.New("Required argument DISCOVERY_URI missing")
)

func Before(c *cli.Context) error {
	var (
		heartbeat time.Duration
		ttl       time.Duration
	)

	if adver := c.String("advertise"); adver == "" {
		return ErrRequireAdvertise
	} else {
		disc.Advertise = adver
	}

	if hbStr := c.String("heartbeat"); hbStr == "" {
		heartbeat = DefaultHeartbeat
	} else {
		if hb, err := time.ParseDuration(hbStr); err != nil {
			log.Warning(err)
			heartbeat = DefaultHeartbeat
		} else {
			heartbeat = hb
		}
	}

	if ttlStr := c.String("ttl"); ttlStr == "" {
		ttl = DefaultTTL
	} else {
		if t, err := time.ParseDuration(ttlStr); err != nil {
			log.Warning(err)
			ttl = DefaultTTL
		} else {
			ttl = t
		}
	}

	if pos := c.Args(); len(pos) != 1 {
		return ErrRequireDiscovery
	} else {
		disc.Discovery = pos[0]
	}

	// register monitor instance
	disc.Register(heartbeat, ttl)

	log.WithFields(log.Fields{"advertise": disc.Advertise, "discovery": disc.Discovery, "heartbeat": heartbeat, "ttl": ttl}).Info("begin advertise")
	return nil
}
