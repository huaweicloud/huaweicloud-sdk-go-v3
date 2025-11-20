package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskStatus struct {

	// **参数解释**： 任务ID，供调用者查询任务进度。获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	JobID *string `json:"jobID,omitempty"`
}

func (o TaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStatus struct{}"
	}

	return strings.Join([]string{"TaskStatus", string(data)}, " ")
}
