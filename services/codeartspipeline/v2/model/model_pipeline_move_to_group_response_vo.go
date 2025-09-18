package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineMoveToGroupResponseVo **参数解释**： 响应详情。 **取值范围**： 不涉及。
type PipelineMoveToGroupResponseVo struct {

	// **参数解释**： 响应码。 **取值范围**： - failed：失败。 - success：成功。
	Code string `json:"code"`

	// **参数解释**： 流水线ID。 **取值范围**： 32位字符，由数字和字母组成。
	PipelineId string `json:"pipeline_id"`

	// **参数解释**： 流水线名。 **取值范围**： 不涉及。
	PipelineName string `json:"pipeline_name"`
}

func (o PipelineMoveToGroupResponseVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineMoveToGroupResponseVo struct{}"
	}

	return strings.Join([]string{"PipelineMoveToGroupResponseVo", string(data)}, " ")
}
