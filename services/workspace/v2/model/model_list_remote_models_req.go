package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemoteModelsReq 查询供应商远程模型列表请求。
type ListRemoteModelsReq struct {

	// 供应商主键ID。传入时，其他空字段从数据库已保存的供应商记录中补充。
	Id *string `json:"id,omitempty"`

	// 供应商类型。
	ProviderType *string `json:"provider_type,omitempty"`

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商API Key（SCC加密存储）。
	ApiKey *string `json:"api_key,omitempty"`

	// 供应商base_url。
	BaseUrl *string `json:"base_url,omitempty"`

	CustomConfig *ProviderCustomConfig `json:"custom_config,omitempty"`

	// 每页数量，默认20，最大100。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRemoteModelsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemoteModelsReq struct{}"
	}

	return strings.Join([]string{"ListRemoteModelsReq", string(data)}, " ")
}
