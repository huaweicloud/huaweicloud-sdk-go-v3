package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetListRequestBody 获取资产列表的请求体
type AssetListRequestBody struct {

	// 资产序列号
	AssetSnList *[]string `json:"asset_sn_list,omitempty"`

	// 资产大类
	AssetType *[]string `json:"asset_type,omitempty"`

	// 首次入库时间搜索起始时间
	StartFirstInboundTime *string `json:"start_first_inbound_time,omitempty"`

	// 首次入库时间搜索结束时间
	EndFirstInboundTime *string `json:"end_first_inbound_time,omitempty"`

	// 资产模型
	Model *string `json:"model,omitempty"`

	// 资产名称
	Name *string `json:"name,omitempty"`

	// 偏移量
	Offset int32 `json:"offset"`

	// 每页条目数量
	Limit int32 `json:"limit"`
}

func (o AssetListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetListRequestBody struct{}"
	}

	return strings.Join([]string{"AssetListRequestBody", string(data)}, " ")
}
