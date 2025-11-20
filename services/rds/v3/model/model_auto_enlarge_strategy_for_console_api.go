package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoEnlargeStrategyForConsoleApi 自动扩容策略
type AutoEnlargeStrategyForConsoleApi struct {

	// 是否开启自动扩容。
	Option *bool `json:"option,omitempty"`

	// 扩容上限，单位GB。“option”为true时，该参数必填。
	LimitSize *int64 `json:"limitSize,omitempty"`

	// 可用存储空间百分比，小于等于此值或者为10GB时触发扩容。“option”为true时，该参数必填。
	TriggerAvailablePercent *int64 `json:"triggerAvailablePercent,omitempty"`

	// 每次自动扩容的步长，单位为百分比，即每次自动扩容当前存储空间的百分比。取值范围为5%~50%。
	StepPercent *int64 `json:"stepPercent,omitempty"`
}

func (o AutoEnlargeStrategyForConsoleApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoEnlargeStrategyForConsoleApi struct{}"
	}

	return strings.Join([]string{"AutoEnlargeStrategyForConsoleApi", string(data)}, " ")
}
