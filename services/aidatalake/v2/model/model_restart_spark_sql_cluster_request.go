package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartSparkSqlClusterRequest Request Object
type RestartSparkSqlClusterRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：端点名称，用于标识SparkSql集群的接入点。 **约束限制**：不涉及。 **取值范围**：只能由小写字母、数字及中划线组成，必须以小写字母开头，以小写字母或数字结尾，且长度为1~63个字符。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	Body *RestartSparkSqlClusterRequestBody `json:"body,omitempty"`
}

func (o RestartSparkSqlClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartSparkSqlClusterRequest struct{}"
	}

	return strings.Join([]string{"RestartSparkSqlClusterRequest", string(data)}, " ")
}
