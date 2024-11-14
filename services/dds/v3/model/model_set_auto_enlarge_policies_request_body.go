package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetAutoEnlargePoliciesRequestBody struct {

	// 自动扩容开关。 - on:开启磁盘自动扩容策略。 - off: 关闭磁盘自动扩容策略。 默认值为on。
	SwitchOption *string `json:"switch_option,omitempty"`

	// 磁盘自动扩容策略。 最大支持设置10个实例的策略。
	Policies []DiskAutoExpansionPolicy `json:"policies"`
}

func (o SetAutoEnlargePoliciesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePoliciesRequestBody struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePoliciesRequestBody", string(data)}, " ")
}
