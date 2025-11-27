package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Toleration 用于定义 Kubernetes Pod对Taint(污点)的容忍规则
type Toleration struct {

	// 表示要容忍的污点键名
	Key *string `json:"key,omitempty"`

	// 定义Key与Value之间的关系，可选值为Exists或Equal，默认为Equal
	Operator *string `json:"operator,omitempty"`

	// 表示要匹配的污点值，当Operator为Exists时，Value应为空
	Value *string `json:"value,omitempty"`

	// 指定要匹配的污点效果，可选值为 NoSchedule、PreferNoSchedule或NoExecute，如果为空，表示匹配所有效果
	Effect *string `json:"effect,omitempty"`

	// 仅对NoExecute效果有效，表示Pod能容忍污点的时间
	TolerationSeconds *int32 `json:"tolerationSeconds,omitempty"`
}

func (o Toleration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Toleration struct{}"
	}

	return strings.Join([]string{"Toleration", string(data)}, " ")
}
