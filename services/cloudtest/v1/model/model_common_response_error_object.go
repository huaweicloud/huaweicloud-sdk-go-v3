package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonResponseErrorObject struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误原因
	Reason *string `json:"reason,omitempty"`
}

func (o CommonResponseErrorObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonResponseErrorObject struct{}"
	}

	return strings.Join([]string{"CommonResponseErrorObject", string(data)}, " ")
}
