package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// indicator create request
type IndicatorCreateRequest struct {

	// 指标名称
	Name *string `json:"name,omitempty"`

	// 版本号
	FormatVersion *int32 `json:"format_version,omitempty"`

	// 类型（SIMULATION,PLAYBOOK,MANUAL,INSTANCE,DATA_SOURCE）
	Type *string `json:"type,omitempty"`

	// 触发标志
	TriggerFlag *bool `json:"trigger_flag,omitempty"`

	DataObject *CreateIndicatorDetail `json:"data_object,omitempty"`
}

func (o IndicatorCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorCreateRequest struct{}"
	}

	return strings.Join([]string{"IndicatorCreateRequest", string(data)}, " ")
}
