package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRouteTableRequest struct {
	Body *CreateRoutetableReqBody `json:"body,omitempty"`
}

func (o CreateRouteTableRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"CreateRouteTableRequest", string(data)}, " ")
}
