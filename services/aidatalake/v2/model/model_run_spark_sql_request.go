package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSparkSqlRequest Request Object
type RunSparkSqlRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：服务事务ID，用于链路追踪和问题定位。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *RunSparkSqlRequestBody `json:"body,omitempty"`
}

func (o RunSparkSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSparkSqlRequest struct{}"
	}

	return strings.Join([]string{"RunSparkSqlRequest", string(data)}, " ")
}
