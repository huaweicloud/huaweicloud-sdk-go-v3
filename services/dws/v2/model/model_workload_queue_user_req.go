package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueUserReq **参数解释**： 资源池用户信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type WorkloadQueueUserReq struct {

	// **参数解释**： 资源池名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	QueueName string `json:"queue_name"`

	// **参数解释**： 资源池用户列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UserList []WorkloadQueueUserReqUserList `json:"user_list"`
}

func (o WorkloadQueueUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueUserReq struct{}"
	}

	return strings.Join([]string{"WorkloadQueueUserReq", string(data)}, " ")
}
