package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDisasterRecoveryRequest Request Object
type CreateDisasterRecoveryRequest struct {
	Body *CreateDisasterRecoveryReq `json:"body,omitempty"`
}

func (o CreateDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryRequest", string(data)}, " ")
}
