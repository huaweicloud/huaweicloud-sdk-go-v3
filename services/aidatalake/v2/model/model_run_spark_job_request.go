package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSparkJobRequest Request Object
type RunSparkJobRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：服务事务ID，用于链路追踪和防止重复提交。 **约束限制**：不涉及。 **取值范围**：采用UUID格式，长度为36个字符。 **默认取值**：不涉及。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *RunSparkJobRequestBody `json:"body,omitempty"`
}

func (o RunSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSparkJobRequest struct{}"
	}

	return strings.Join([]string{"RunSparkJobRequest", string(data)}, " ")
}
