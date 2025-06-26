package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageRetention struct {

	// 老化策略ID
	Id int32 `json:"id"`

	// 算法
	Algorithm string `json:"algorithm"`

	// 匹配规则
	Rules []RetentionRuleResponseBody `json:"rules"`

	Trigger *TriggerConfig `json:"trigger"`

	// 是否开启策略
	Enabled bool `json:"enabled"`

	// 策略名称
	Name string `json:"name"`

	// 命名空间ID
	NamespaceId int32 `json:"namespace_id"`

	// 命名空间名
	NamespaceName string `json:"namespace_name"`
}

func (o ImageRetention) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageRetention struct{}"
	}

	return strings.Join([]string{"ImageRetention", string(data)}, " ")
}
