package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDisasterRecoveryRequest struct {
	Body *CreateDisasterRecovery `json:"body,omitempty"`
}

func (o CreateDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryRequest", string(data)}, " ")
}
