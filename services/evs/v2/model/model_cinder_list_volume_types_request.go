package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CinderListVolumeTypesRequest struct {
}

func (o CinderListVolumeTypesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CinderListVolumeTypesRequest struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTypesRequest", string(data)}, " ")
}
