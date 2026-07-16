package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigtemplatesSpec **参数解释**：节点配置模板规格。
type NodeconfigtemplatesSpec struct {

	// **参数解释**：节点配置模板列表。
	Templates *[]NodeconfigTemplateItem `json:"templates,omitempty"`
}

func (o NodeconfigtemplatesSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigtemplatesSpec struct{}"
	}

	return strings.Join([]string{"NodeconfigtemplatesSpec", string(data)}, " ")
}
