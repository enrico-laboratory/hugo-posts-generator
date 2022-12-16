package hugopostgenerator

import (
	"embed"
	"github.com/rs/zerolog"
	"os"
	"time"
)

//go:embed templates
var fs embed.FS
var logger zerolog.Logger

func NewClient() {

	logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).With().Timestamp().Caller().Logger()

	//notionClient, err := notionclient.NewClient()
	//if err != nil {
	//	logger.Error().Err(err).Send()
	//}

}
