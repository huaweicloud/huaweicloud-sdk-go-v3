package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransJobRejectBody struct {

	// 拒绝理由
	Reason *string `json:"reason,omitempty"`
}

func (o TransJobRejectBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransJobRejectBody struct{}"
	}

	return strings.Join([]string{"TransJobRejectBody", string(data)}, " ")
}
