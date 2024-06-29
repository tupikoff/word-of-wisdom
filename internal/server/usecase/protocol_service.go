package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/tupikoff/word-of-wisdom/internal/server/domain"
	"github.com/tupikoff/word-of-wisdom/pkg/hashcash"
	"github.com/tupikoff/word-of-wisdom/pkg/random"
)

const (
	defaultDifficulty            = 20
	defaultChallengeStringLength = 30
)

// ProtocolService are Request-Challenge Protocol
type ProtocolService struct {
	wisdomRepository   wisdomRepository
	registerRepository registerRepository
}

func NewProtocolService(
	wisdomRepository wisdomRepository,
	registerRepository registerRepository,
) *ProtocolService {
	return &ProtocolService{wisdomRepository: wisdomRepository, registerRepository: registerRepository}
}

func (c ProtocolService) Execute(
	ctx context.Context,
	request string,
) (response string, err error) {
	requestData := strings.Split(request, "|")

	switch requestData[0] {
	case "request":
		randString := random.String(defaultChallengeStringLength)
		record := domain.RegisterRecord{
			RandString: randString,
			HashString: base64.StdEncoding.EncodeToString([]byte(randString)),
			Difficulty: defaultDifficulty,
		}
		err := c.registerRepository.Save(ctx, record)
		if err != nil {
			return "", err
		}
		log.Printf("saved record: %+v", record)
		response = fmt.Sprintf("challenge|%s:%d", record.RandString, record.Difficulty)
	case "response":
		if len(requestData) > 1 {
			solution := requestData[1]
			hc, err := hashcash.NewFromString(solution)
			if err != nil {
				return "", fmt.Errorf("%w: %s", domain.ErrHashReadError, err.Error())
			}

			rec, err := c.registerRepository.Get(ctx, hc.Rand)
			if err != nil {
				return "", err
			}

			if rec.Difficulty != hc.Bits {
				return "", domain.ErrDifficultyNotMatchWithRegistered
			}

			if !hc.IsHashValid() {
				return "", domain.ErrHashNotValid
			}
			response = fmt.Sprintf("granted|%s", c.wisdomRepository.Read())
		}
	default:
		return "", domain.ErrUnknownCommand
	}

	return response, nil
}
