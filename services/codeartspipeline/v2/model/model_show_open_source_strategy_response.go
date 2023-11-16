package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOpenSourceStrategyResponse Response Object
type ShowOpenSourceStrategyResponse struct {

	// 策略ID
	Id *string `json:"id,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 策略级别
	Level *string `json:"level,omitempty"`

	// 父策略ID
	ParentId *string `json:"parent_id,omitempty"`

	// 策略版本
	Version *string `json:"version,omitempty"`

	// 是否启用
	IsValid *bool `json:"is_valid,omitempty"`

	// 是否系统级策略
	IsPublic *bool `json:"is_public,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新人
	Updater *string `json:"updater,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	Content        *OpenSourceRuleContent `json:"content,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowOpenSourceStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpenSourceStrategyResponse struct{}"
	}

	return strings.Join([]string{"ShowOpenSourceStrategyResponse", string(data)}, " ")
}
