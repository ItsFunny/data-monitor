package constants

import (
	"fmt"
	"github.com/pkg/errors"
)

// pcap
var (
	PCAP_OPENLIVE_ERROR = errors.New("pcap#openlive occur error")
)

func MissingError(key, value string) error {
	return errors.New(fmt.Sprintf("missing %v of %v ", value, key))
}
