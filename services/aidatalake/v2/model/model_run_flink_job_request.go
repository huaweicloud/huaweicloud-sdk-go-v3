package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunFlinkJobRequest Request Object
type RunFlinkJobRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：幂等性校验唯一标识，非必填。1小时内携带相同token的请求，将视为重复请求，返回上一次成功的响应结果。 **约束限制**：长度0-128。 **取值范围**：不涉及。 **默认取值**：不涉及。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *RunFlinkJobRequestBody `json:"body,omitempty"`
}

func (o RunFlinkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFlinkJobRequest struct{}"
	}

	return strings.Join([]string{"RunFlinkJobRequest", string(data)}, " ")
}
