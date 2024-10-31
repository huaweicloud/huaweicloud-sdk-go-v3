package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetAutoPolicyRequestBody struct {

	// 设置磁盘自动扩容的实例组ID。
	InstanceIds []string `json:"instance_ids"`

	// 自动扩容开关。  “on”，表示开启磁盘自动扩容策略。  “off”，表示关闭磁盘自动扩容策略。 默认值为“on”。
	SwitchOption *string `json:"switch_option,omitempty"`

	// 磁盘自动扩容策略
	Policy []DiskAutoExpansionPolicy `json:"policy"`
}

func (o SetAutoPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetAutoPolicyRequestBody", string(data)}, " ")
}
