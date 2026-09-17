package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOperationalTaskRequest Request Object
type UpdateOperationalTaskRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 任务ID，可先通过任务列表接口查询所有任务。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskId string `json:"task_id"`

	Body *TaskInfo `json:"body,omitempty"`
}

func (o UpdateOperationalTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOperationalTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateOperationalTaskRequest", string(data)}, " ")
}
