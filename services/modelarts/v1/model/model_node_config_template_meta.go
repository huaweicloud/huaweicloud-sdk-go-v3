package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeConfigTemplateMeta 节点配置模板的metadata信息。
type NodeConfigTemplateMeta struct {

	// **参数解释**： 节点配置模板的名称。 **取值范围**： 不涉及。
	Name string `json:"name"`
}

func (o NodeConfigTemplateMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeConfigTemplateMeta struct{}"
	}

	return strings.Join([]string{"NodeConfigTemplateMeta", string(data)}, " ")
}
