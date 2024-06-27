package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonResponseErrorPageResultBasicAwInfo struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Detail *PageResultBasicAwInfo `json:"detail,omitempty"`

	// 错误原因
	Reason *string `json:"reason,omitempty"`
}

func (o CommonResponseErrorPageResultBasicAwInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonResponseErrorPageResultBasicAwInfo struct{}"
	}

	return strings.Join([]string{"CommonResponseErrorPageResultBasicAwInfo", string(data)}, " ")
}
