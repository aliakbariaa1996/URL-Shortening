package shorteningURL

import (
	"context"
)

type Store interface {
	SaveUrlMapping(ctx context.Context, shortUrl string, originalUrl string) (string, error)
	RetrieveInitialUrl(ctx context.Context, shortUrl string) (string, error)
}
