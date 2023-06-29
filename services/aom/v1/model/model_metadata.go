package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metadata 任务节点元数据。
type Metadata struct {

	// 节点类型。
	Type *string `json:"type,omitempty"`

	// 配置信息。
	Configuration map[string]interface{} `json:"configuration,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
