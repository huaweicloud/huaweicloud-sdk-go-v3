package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAvailableAzResponsebodyAvailabilityZone struct {

	// 可用区ID
	Id *string `json:"id,omitempty"`

	// 显示名称
	DisplayName *string `json:"display_name,omitempty"`

	Locales *ShowAvailableAzResponsebodyLocales `json:"locales,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 局点ID
	RegionId *string `json:"region_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`
}

func (o ShowAvailableAzResponsebodyAvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableAzResponsebodyAvailabilityZone struct{}"
	}

	return strings.Join([]string{"ShowAvailableAzResponsebodyAvailabilityZone", string(data)}, " ")
}
