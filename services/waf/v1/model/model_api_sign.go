package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiSign API签名配置
type ApiSign struct {

	// API签名状态（0：禁用，1：启用）
	Status *int32 `json:"status,omitempty"`
}

func (o ApiSign) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiSign struct{}"
	}

	return strings.Join([]string{"ApiSign", string(data)}, " ")
}
