package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCustomfieldsRequest struct {
	// 项目ID

	ProjectId string `json:"project_id"`

	Body *CreateCustomfieldV1Req `json:"body,omitempty"`
}

func (o CreateCustomfieldsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCustomfieldsRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomfieldsRequest", string(data)}, " ")
}
