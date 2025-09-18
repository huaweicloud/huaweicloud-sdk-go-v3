package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineBasicInfo struct {

	// **参数解释**： 项目id。 **取值范围**： 32位字符，由数字和字母组成。
	ProjectId string `json:"project_id"`

	// **参数解释**： 项目名称。 **取值范围**： 不涉及。
	ProjectName string `json:"project_name"`

	// **参数解释**： 流水线id。 **取值范围**： 32位字符，由数字和字母组成。
	PipelineId string `json:"pipeline_id"`

	// **参数解释**： 流水线名称。 **取值范围**： 不涉及。
	PipelineName string `json:"pipeline_name"`

	// **参数解释**： 流水线创建人id。 **取值范围**： 32位字符，由数字和字母组成。
	CreatorId string `json:"creator_id"`

	// **参数解释**： 流水线创建人名字。 **取值范围**： 不涉及。
	CreatorName string `json:"creator_name"`

	// **参数解释**： 流水线执行人id。 **取值范围**： 32位字符，由数字和字母组成。
	ExecutorId string `json:"executor_id"`

	// **参数解释**： 流水线执行人名字。 **取值范围**： 不涉及。
	ExecutorName string `json:"executor_name"`

	// **参数解释**： 启动时间。 **取值范围**： 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreateTime string `json:"create_time"`

	// **参数解释**： 用户是否关注流水线。 **取值范围**： - true：关注流水线。 - false：未关注流水线。
	Watched string `json:"watched"`
}

func (o PipelineBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineBasicInfo struct{}"
	}

	return strings.Join([]string{"PipelineBasicInfo", string(data)}, " ")
}
