package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().
		Str("version", SemVer).
		Str("gitCommit", GitCommit).
		Msg("APP starting ...")
	houseLoan()
}
