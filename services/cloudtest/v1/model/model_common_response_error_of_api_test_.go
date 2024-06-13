package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonResponseErrorOfApiTest struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误原因
	Reason *string `json:"reason,omitempty"`
}

func (o CommonResponseErrorOfApiTest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonResponseErrorOfApiTest struct{}"
	}

	return strings.Join([]string{"CommonResponseErrorOfApiTest", string(data)}, " ")
}
