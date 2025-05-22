package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterRequestBody **参数解释**： 删除集群请求信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type DeleteClusterRequestBody struct {

	// **参数解释**： 集群需要保留的快照数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	KeepLastManualSnapshot int32 `json:"keep_last_manual_snapshot"`
}

func (o DeleteClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequestBody", string(data)}, " ")
}
