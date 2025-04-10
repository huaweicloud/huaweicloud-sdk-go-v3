package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchApproveRequestBody struct {

	// 通过/拒绝原因
	Reason *string `json:"reason,omitempty"`

	// 工单id列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o BatchApproveRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchApproveRequestBody struct{}"
	}

	return strings.Join([]string{"BatchApproveRequestBody", string(data)}, " ")
}
