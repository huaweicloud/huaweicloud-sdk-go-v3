package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobDetailsRequest Request Object
type ListJobDetailsRequest struct {

	// **参数解释**： 任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	JobId string `json:"job_id"`
}

func (o ListJobDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListJobDetailsRequest", string(data)}, " ")
}
