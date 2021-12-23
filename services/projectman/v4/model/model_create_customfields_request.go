package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCustomfieldsRequest struct {
	// devcloud的项目ID

	ProjectId string `json:"project_id"`

	Body *CreateCustomfieldV1Req `json:"body,omitempty"`
}

func (o CreateCustomfieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomfieldsRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomfieldsRequest", string(data)}, " ")
}
