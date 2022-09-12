package shorteningURL

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

type ShortingURL struct{}

func (s *UseCase) GenerateShortLink(initialLink string, userId string) (string, error) {
	urlHashBytes, err := s.sha256Of(initialLink + userId)
	if err != nil {
		s.logger.Errorf(err.Error())
		return "", err
	}
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := s.base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		s.logger.Errorf(err.Error())
		return "", err
	}
	return finalString[:8], nil
}

func (s *UseCase) SaveUrlMapping(ctx context.Context, shortUrl string, longUrl string, userId string) (string, error) {
	res, err := s.shRepo.SaveUrlMapping(ctx, shortUrl, longUrl)
	if err != nil {
		s.logger.Errorf(err.Error())
		return "", err
	}
	return res, err
}
func (s *UseCase) RetrieveInitialUrl(ctx context.Context, shortUrl string) (string, error) {
	res, err := s.shRepo.RetrieveInitialUrl(ctx, shortUrl)
	if err != nil {
		s.logger.Errorf(err.Error())
		return "", err
	}
	return res, err
}

func (s *UseCase) sha256Of(input string) ([]byte, error) {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil), nil
}

func (s *UseCase) base58Encoded(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		s.logger.Errorf(err.Error())
		os.Exit(1)
	}
	return string(encoded), nil
}
