package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ValidateConsumedMessageRequest struct {

	// engine
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResendReq `json:"body,omitempty"`
}

func (o ValidateConsumedMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConsumedMessageRequest struct{}"
	}

	return strings.Join([]string{"ValidateConsumedMessageRequest", string(data)}, " ")
}
