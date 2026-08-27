package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelGroupProviderDetailResp 模型分组供应商关联列表详情项。
type ModelGroupProviderDetailResp struct {

	// 关联记录id。
	Id *string `json:"id,omitempty"`

	// 分组id。
	GroupId *string `json:"group_id,omitempty"`

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商名称。
	ProviderName *string `json:"provider_name,omitempty"`

	// 供应商类型。
	ProviderType *string `json:"provider_type,omitempty"`

	// 供应商base_url。
	BaseUrl *string `json:"base_url,omitempty"`

	// 连接状态。
	ConnectionStatus *string `json:"connection_status,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ModelGroupProviderDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelGroupProviderDetailResp struct{}"
	}

	return strings.Join([]string{"ModelGroupProviderDetailResp", string(data)}, " ")
}
