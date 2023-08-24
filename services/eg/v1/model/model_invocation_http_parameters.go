package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InvocationHttpParameters struct {

	// 对象列表
	HeaderParameters *[]HeaderParameter `json:"header_parameters,omitempty"`
}

func (o InvocationHttpParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvocationHttpParameters struct{}"
	}

	return strings.Join([]string{"InvocationHttpParameters", string(data)}, " ")
}
