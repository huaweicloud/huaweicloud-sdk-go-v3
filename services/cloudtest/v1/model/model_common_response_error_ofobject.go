package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonResponseErrorOfobject struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误原因
	Reason *string `json:"reason,omitempty"`
}

func (o CommonResponseErrorOfobject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonResponseErrorOfobject struct{}"
	}

	return strings.Join([]string{"CommonResponseErrorOfobject", string(data)}, " ")
}
