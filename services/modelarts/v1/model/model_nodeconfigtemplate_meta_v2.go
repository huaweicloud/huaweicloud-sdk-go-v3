package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigtemplateMetaV2 **参数解释**：节点配置模板元数据。
type NodeconfigtemplateMetaV2 struct {

	// **参数解释**：节点名称。 **取值范围**：固定为node-config-template。
	Name string `json:"name"`
}

func (o NodeconfigtemplateMetaV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigtemplateMetaV2 struct{}"
	}

	return strings.Join([]string{"NodeconfigtemplateMetaV2", string(data)}, " ")
}
