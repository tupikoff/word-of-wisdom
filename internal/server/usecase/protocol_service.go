package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/tupikoff/word-of-wisdom/internal/server/domain"
	"github.com/tupikoff/word-of-wisdom/pkg/hashcash"
	"github.com/tupikoff/word-of-wisdom/pkg/random"
)

const (
	defaultDifficulty            = 22
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
		response = fmt.Sprintf("challenge|%s", random.String(defaultChallengeStringLength))
	case "response":
		if len(requestData) > 1 {
			solution := requestData[1]
			hc, err := hashcash.NewFromString(solution)
			if err != nil {
				return "", fmt.Errorf("%w: %s", domain.ErrHashReadError, err.Error())
			}
			if hc.Bits != defaultDifficulty {
				return "", domain.ErrDifficultyNotMatch
			}
			if !hc.IsHashValid() {
				return "", domain.ErrHashNotValid
			}

			record := domain.RegisterRecord{
				HashString: hc.Rand,
			}
			err = c.registerRepository.Save(ctx, record)
			if err != nil {
				if errors.Is(err, domain.ErrRecordAlreadyExists) {
					return "", fmt.Errorf("hash was used: %w", err)
				}
				return "", err
			}
			log.Printf("saved record: %+v", record)

			response = fmt.Sprintf("granted|%s", c.wisdomRepository.Read())
		}
	default:
		return "", domain.ErrUnknownCommand
	}

	return response, nil
}
