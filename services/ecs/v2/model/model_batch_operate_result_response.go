package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchOperateResultResponse struct {
	Status *string `json:"status,omitempty"`

	ServerId *string `json:"server_id,omitempty"`

	FailReason *string `json:"fail_reason,omitempty"`
}

func (o BatchOperateResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateResultResponse struct{}"
	}

	return strings.Join([]string{"BatchOperateResultResponse", string(data)}, " ")
}
