package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkspaceQuotasResponse quotas属性列表
type WorkspaceQuotasResponse struct {

	// 配额允许设置的最大值。
	MaxQuota *int32 `json:"max_quota,omitempty"`

	// 最后修改时间，UTC。如用户未修改过该资源配额,则该值默认为该工作空间的创建时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 资源的唯一标识。
	Resource *string `json:"resource,omitempty"`

	// 当前配额值。配额值为-1代表不限制配额。
	Quota *int32 `json:"quota,omitempty"`

	// 配额允许设置的最小值。
	MinQuota *int32 `json:"min_quota,omitempty"`

	// 配额名称[(中文)](tag:hc,hk)。
	NameCn *string `json:"name_cn,omitempty"`

	// 数量单位[(中文)](tag:hc,hk)。
	UnitCn *string `json:"unit_cn,omitempty"`

	// 工作空间ID，系统生成的32位UUID，不带橫线。默认的工作空间id为'0'。
	NameEn *string `json:"name_en,omitempty"`

	// 数量单位(英文)。
	UnitEn *string `json:"unit_en,omitempty"`

	// 已用配额值。当quota为-1（不限制配额）时，used_quota为null。
	UsedQuota *int32 `json:"used_quota,omitempty"`
}

func (o WorkspaceQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkspaceQuotasResponse struct{}"
	}

	return strings.Join([]string{"WorkspaceQuotasResponse", string(data)}, " ")
}
