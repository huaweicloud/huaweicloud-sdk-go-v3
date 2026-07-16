package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServicePodResponse struct {

	// **参数解释：** od ID。 **取值范围：** 不涉及。
	PodId *string `json:"pod_id,omitempty"`

	// **参数解释：** pod名字。 **取值范围：** 不涉及。
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释：** pod所在node的IP。 **取值范围：** 不涉及。
	PodNodeIp *string `json:"pod_node_ip,omitempty"`

	// **参数解释：** pod所在node的名字。 **取值范围：** 不涉及。
	PodNodeName *string `json:"pod_node_name,omitempty"`

	// **参数解释：** pod角色。 **取值范围：** 不涉及。
	PodRole *string `json:"pod_role,omitempty"`

	// **参数解释：** pod服务状态。 **取值范围：** 有7种状态。RUNNING（运行中）、PENDING（未就绪）、SUCCEEDED（成功）、FAILED（失败）、ABNORMAL（异常）、UNKNOWN（未知）、DELETED（已删除）。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 最近更新时间。 **取值范围：** 不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o ServicePodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServicePodResponse struct{}"
	}

	return strings.Join([]string{"ServicePodResponse", string(data)}, " ")
}
