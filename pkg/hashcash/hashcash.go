package hashcash

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tupikoff/word-of-wisdom/pkg/random"
)

type hashCash struct {
	Ver      int    // Hashcash format version, 1 (which supersedes version 0).
	Bits     int    // Number of "partial pre-image" (zero) bits in the hashed code.
	Date     string // The time that the message was sent, in the format YYMMDD[hhmm[ss]].
	Resource string // Resource data string being transmitted, e.g., an IP address or email address.
	Ext      string // Extension (optional; ignored in version 1).
	Rand     string // String of random characters, encoded in base-64 format.
	Counter  string // Binary counter, encoded in base-64 format.

	counter int // internal counter
}

func New(resource, rand string, bits int) *hashCash {
	c := random.Int()
	hc := &hashCash{
		Ver:      1,
		Bits:     bits,
		Date:     time.Now().Format("0601021504"),
		Resource: resource,
		Ext:      "",
		Rand:     Hash(rand),
		Counter:  Hash(strconv.Itoa(c)),
		counter:  c,
	}
	hc.calculate()
	return hc
}

func NewFromString(s string) (*hashCash, error) {
	ss := strings.Split(s, ":")
	if len(ss) != 7 {
		return nil, errors.New("invalid hash string, parameters must be 7")
	}
	v, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil, err
	}
	b, err := strconv.Atoi(ss[1])
	if err != nil {
		return nil, err
	}
	return &hashCash{
		Ver:      v,
		Bits:     b,
		Date:     ss[2],
		Resource: ss[3],
		Ext:      ss[4],
		Rand:     ss[5],
		Counter:  ss[6],
	}, nil
}

func (h *hashCash) String() string {
	return fmt.Sprintf("%d:%d:%s:%s:%s:%s:%s",
		h.Ver, h.Bits, h.Date, h.Resource, h.Ext, h.Rand, h.Counter)
}

func (h *hashCash) Hash() hash {
	return sha1.Sum([]byte(h.String()))
}

func (h *hashCash) IsHashValid() bool {
	return h.Hash().IsValid(h.Bits)
}

func (h *hashCash) calculate() {
	for {
		if h.IsHashValid() {
			return
		}
		h.counter++
		h.Counter = base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(h.counter)))
	}
}

func Hash(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
