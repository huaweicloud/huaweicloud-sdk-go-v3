package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelGroupProviderItemResp 模型分组关联供应商响应项。
type ModelGroupProviderItemResp struct {

	// 关联记录id。
	Id *string `json:"id,omitempty"`

	// 模型组id。
	GroupId *string `json:"group_id,omitempty"`

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ModelGroupProviderItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelGroupProviderItemResp struct{}"
	}

	return strings.Join([]string{"ModelGroupProviderItemResp", string(data)}, " ")
}
