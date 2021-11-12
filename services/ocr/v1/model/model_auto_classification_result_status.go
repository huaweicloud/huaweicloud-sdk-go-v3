package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AutoClassificationResultStatus struct {
	// 指示各对应票证的状态码

	ErrorCode *string `json:"error_code,omitempty"`
	// 指示各对应票证的状态信息

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o AutoClassificationResultStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoClassificationResultStatus struct{}"
	}

	return strings.Join([]string{"AutoClassificationResultStatus", string(data)}, " ")
}
