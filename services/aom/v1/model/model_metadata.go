package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务节点形状信息。
type Metadata struct {

	// 形状,节点类型。
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
