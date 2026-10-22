package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiInfo API信息
type ApiInfo struct {

	// API签名状态（0：禁用，1：启用）
	Status *int32 `json:"status,omitempty"`

	// 访问密钥ID
	SourceId *string `json:"source_id,omitempty"`

	// 秘密访问密钥
	SourceSecret *string `json:"source_secret,omitempty"`
}

func (o ApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiInfo struct{}"
	}

	return strings.Join([]string{"ApiInfo", string(data)}, " ")
}
