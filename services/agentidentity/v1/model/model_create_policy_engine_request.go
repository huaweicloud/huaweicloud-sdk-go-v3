package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyEngineRequest Request Object
type CreatePolicyEngineRequest struct {
	Body *CreatePolicyEngineReqBody `json:"body,omitempty"`
}

func (o CreatePolicyEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyEngineRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyEngineRequest", string(data)}, " ")
}
