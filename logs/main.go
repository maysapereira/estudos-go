package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Int("EnployeeID", 1001).
		Msg("Getting employee information")

	log.Debug().
		Str("Name", "Maysa").
		Send()
}
