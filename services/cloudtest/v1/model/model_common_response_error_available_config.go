package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonResponseErrorAvailableConfig struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Detail *AvailableConfig `json:"detail,omitempty"`

	// 错误原因
	Reason *string `json:"reason,omitempty"`
}

func (o CommonResponseErrorAvailableConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonResponseErrorAvailableConfig struct{}"
	}

	return strings.Join([]string{"CommonResponseErrorAvailableConfig", string(data)}, " ")
}
