package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDirectoryByResultCommitIdRequest Request Object
type ListDirectoryByResultCommitIdRequest struct {

	// **参数解释**： 演化任务标识符。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	EvolveTaskId string `json:"evolve_task_id"`

	// **参数解释**： 演化任务结果的commit_id。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-256]个字符。 **默认取值**： 不涉及
	CommitId string `json:"commit_id"`

	// **参数解释**： 从哪个轮次开始查询。 **约束限制**： 不涉及 **取值范围**： [-1-10000]。 **默认取值**： 不涉及
	Iteration int32 `json:"iteration"`
}

func (o ListDirectoryByResultCommitIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectoryByResultCommitIdRequest struct{}"
	}

	return strings.Join([]string{"ListDirectoryByResultCommitIdRequest", string(data)}, " ")
}
