package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetManagerPasswordRequest Request Object
type ResetManagerPasswordRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResetManagerPasswordReq `json:"body,omitempty"`
}

func (o ResetManagerPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetManagerPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetManagerPasswordRequest", string(data)}, " ")
}
