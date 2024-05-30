package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalInfoParam struct {

	// 审批单ID列表，填写String类型替代Long类型。
	Ids []string `json:"ids"`

	// 审批单信息，审批人填写的审批意见。
	Msg string `json:"msg"`
}

func (o ApprovalInfoParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalInfoParam struct{}"
	}

	return strings.Join([]string{"ApprovalInfoParam", string(data)}, " ")
}
