package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteProvidersReq 批量解绑供应商从模型组请求。
type BatchDeleteProvidersReq struct {

	// 供应商id列表。
	ProviderIds []string `json:"provider_ids"`
}

func (o BatchDeleteProvidersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProvidersReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteProvidersReq", string(data)}, " ")
}
