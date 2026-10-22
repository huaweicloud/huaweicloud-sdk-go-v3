package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuraDataframeJobRequest Request Object
type ShowAuraDataframeJobRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：作业ID。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	JobId string `json:"job_id"`
}

func (o ShowAuraDataframeJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuraDataframeJobRequest struct{}"
	}

	return strings.Join([]string{"ShowAuraDataframeJobRequest", string(data)}, " ")
}
