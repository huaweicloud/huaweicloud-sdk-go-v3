package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEipRecordSetRequest struct {
	Region       string        `json:"region"`
	FloatingipId string        `json:"floatingip_id"`
	Body         *CreatePtrReq `json:"body,omitempty"`
}

func (o CreateEipRecordSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEipRecordSetRequest struct{}"
	}

	return strings.Join([]string{"CreateEipRecordSetRequest", string(data)}, " ")
}
