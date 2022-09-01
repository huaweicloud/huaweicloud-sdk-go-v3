package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 失败原因对象。
type FailReason struct {

	// 错误码
	FailCode *string `json:"fail_code,omitempty" xml:"fail_code"`

	// 边缘云失败原因列表。包含所边缘云的失败原因。
	FailMessage *string `json:"fail_message,omitempty" xml:"fail_message"`
}

func (o FailReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailReason struct{}"
	}

	return strings.Join([]string{"FailReason", string(data)}, " ")
}
