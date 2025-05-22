package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLogicalClusterRequestBody **参数解释**： 切换逻辑集群开关请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type EnableLogicalClusterRequestBody struct {

	// **参数解释**： 切换逻辑集群开关请求。 **约束限制**： 不涉及。 **取值范围**： true：打开开关 false：关闭开关 **默认取值**： 不涉及。
	Enable bool `json:"enable"`
}

func (o EnableLogicalClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLogicalClusterRequestBody struct{}"
	}

	return strings.Join([]string{"EnableLogicalClusterRequestBody", string(data)}, " ")
}
