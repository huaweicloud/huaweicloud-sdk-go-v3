package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigTemplateItem **参数解释**：节点自定义配置模板
type NodeconfigTemplateItem struct {

	// **参数解释**：支持的配置项列表和默认值。
	Configs *interface{} `json:"configs,omitempty"`
}

func (o NodeconfigTemplateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigTemplateItem struct{}"
	}

	return strings.Join([]string{"NodeconfigTemplateItem", string(data)}, " ")
}
