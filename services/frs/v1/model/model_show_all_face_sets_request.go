package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAllFaceSetsRequest struct {
}

func (o ShowAllFaceSetsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAllFaceSetsRequest struct{}"
	}

	return strings.Join([]string{"ShowAllFaceSetsRequest", string(data)}, " ")
}
