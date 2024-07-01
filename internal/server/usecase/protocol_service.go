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
	"github.com/tupikoff/word-of-wisdom/pkg/tcp"
)

// ProtocolService are Request-Challenge Protocol https://en.wikipedia.org/wiki/Proof_of_work
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
	connection *tcp.Connection,
) (err error) {
	// 1. request service
	request, err := connection.Read()
	if err != nil {
		return fmt.Errorf("protocol execute: read: %v", err)
	}
	requestData := strings.Split(request, "|")
	if requestData[0] != "request" {
		return errors.New("protocol execute: expected `request`")
	}

	// 2. choose
	randString := random.String(random.IntIn(20, 40))
	difficulty := random.IntIn(19, 23)
	// 3. challenge
	response := fmt.Sprintf("challenge|%s:%d", randString, difficulty)
	err = connection.Write(response)
	if err != nil {
		return fmt.Errorf("protocol execute: write: %v", err)
	}

	// 5. response
	request, err = connection.Read()
	if err != nil {
		return fmt.Errorf("protocol execute: read: %v", err)
	}
	requestData = strings.Split(request, "|")
	if requestData[0] != "response" {
		return errors.New("protocol execute: expected `response`")
	}
	// 6. verify
	if len(requestData) != 2 {
		return errors.New("protocol execute: expected `response` len 2")
	}
	solution := requestData[1]
	hc, err := hashcash.NewFromString(solution)
	if err != nil {
		return fmt.Errorf("%w: %s", domain.ErrHashReadError, err.Error())
	}
	if hc.Rand != hashcash.Hash(randString) {
		return domain.ErrRandomStringNotMatch
	}
	if hc.Bits != difficulty {
		return domain.ErrDifficultyNotMatch
	}
	if !hc.IsHashValid() {
		return domain.ErrHashNotValid
	}
	record := domain.RegisterRecord{
		HashString: hc.Rand,
		Difficulty: difficulty,
	}
	err = c.registerRepository.Save(ctx, record)
	if err != nil {
		if errors.Is(err, domain.ErrRecordAlreadyExists) {
			return fmt.Errorf("hash was used: %w", err)
		}
		return err
	}
	log.Printf("saved record: %+v", record)

	// 7. grant service
	response = fmt.Sprintf("granted|%s", c.wisdomRepository.Read())
	err = connection.Write(response)
	if err != nil {
		return fmt.Errorf("protocol execute: write: %v", err)
	}

	return nil
}
