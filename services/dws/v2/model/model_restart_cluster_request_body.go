package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartClusterRequestBody **参数解释**： 重启集群的请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RestartClusterRequestBody struct {

	// **参数解释**： 重启标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Restart *interface{} `json:"restart"`
}

func (o RestartClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartClusterRequestBody struct{}"
	}

	return strings.Join([]string{"RestartClusterRequestBody", string(data)}, " ")
}
