package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationalTaskIdListRequest struct {

	// **参数解释**： 任务ID列表。 **约束限制**： 不涉及。 **取值范围**： 非空。 **默认取值**： null
	TaskIds *[]string `json:"task_ids,omitempty"`
}

func (o OperationalTaskIdListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationalTaskIdListRequest struct{}"
	}

	return strings.Join([]string{"OperationalTaskIdListRequest", string(data)}, " ")
}
