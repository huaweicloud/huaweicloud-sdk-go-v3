package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEnhancedConnectionRoutesRequestBody struct {

	// 路由名称，长度限制：1-64个字符。
	Name string `json:"name"`

	// 路由网段范围。
	Cidr string `json:"cidr"`
}

func (o CreateEnhancedConnectionRoutesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnhancedConnectionRoutesRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEnhancedConnectionRoutesRequestBody", string(data)}, " ")
}
