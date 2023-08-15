package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalInfoParam struct {

	// 审批单id列表
	Ids []int64 `json:"ids"`

	// 审批单信息
	Msg string `json:"msg"`
}

func (o ApprovalInfoParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalInfoParam struct{}"
	}

	return strings.Join([]string{"ApprovalInfoParam", string(data)}, " ")
}
