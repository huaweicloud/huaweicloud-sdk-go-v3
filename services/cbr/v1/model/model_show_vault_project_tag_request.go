package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVaultProjectTagRequest struct {
}

func (o ShowVaultProjectTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVaultProjectTagRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultProjectTagRequest", string(data)}, " ")
}
