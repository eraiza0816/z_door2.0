package pasori

import (
	"fmt"

	"github.com/bamchoh/pasori"
)

func GetID(VID uint16, PID uint16) (string, error) {
	idm, err := pasori.GetID(VID, PID)
	if err != nil {
		return "", fmt.Errorf("failed to get ID: %w", err)
	}

	return idm, nil
}
