package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBindProvidersReq 批量绑定供应商到模型组请求。
type BatchBindProvidersReq struct {

	// 供应商id列表。
	ProviderIds []string `json:"provider_ids"`
}

func (o BatchBindProvidersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindProvidersReq struct{}"
	}

	return strings.Join([]string{"BatchBindProvidersReq", string(data)}, " ")
}
