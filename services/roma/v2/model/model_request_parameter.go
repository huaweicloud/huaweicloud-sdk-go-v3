package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RequestParameter struct {

	// 请求参数名
	ParameterName *string `json:"parameter_name,omitempty"`

	// 请求参数值
	ParameterValue *string `json:"parameter_value,omitempty"`
}

func (o RequestParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestParameter struct{}"
	}

	return strings.Join([]string{"RequestParameter", string(data)}, " ")
}
