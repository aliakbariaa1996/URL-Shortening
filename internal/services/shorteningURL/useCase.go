package shorteningURL

import (
	"context"
	"github.com/aliakbariaa1996/URL-Shortening/config"
	"github.com/aliakbariaa1996/URL-Shortening/internal/services/shorteningURL/store"
	loggerx "github.com/sirupsen/logrus"
)

type UseCase struct {
	shRepo *store.Store
	cfg    *config.Config
	logger *loggerx.Logger
}

func NewShorteningCase(cfg *config.Config, logger *loggerx.Logger, store *store.Store) *UseCase {
	return &UseCase{
		shRepo: store,
		cfg:    cfg,
		logger: logger,
	}
}

type UseService interface {
	GenerateShortLink(initialLink string, userId string) (string, error)
	SaveUrlMapping(ctx context.Context, shortUrl string, longUrl string, userId string) (string, error)
	RetrieveInitialUrl(ctx context.Context, shortUrl string) (string, error)
	sha256Of(input string) ([]byte, error)
	base58Encoded(bytes []byte) (string, error)
}
