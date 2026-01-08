package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberStatus struct {

	// **参数解释**：监听器ID  **取值范围**：不涉及
	ListenerId string `json:"listener_id"`

	// **参数解释**：后端服务器的健康状态。  **取值范围**： - ONLINE：后端服务器正常。 - NO_MONITOR：后端服务器所在的服务器组没有健康检查器。 - OFFLINE：后端服务器关联的ECS服务器不存在或已关机。 - INITIAL：后端云服务器健康检查打开时的初始状态。 - UNKNOWN: 后端云服务器组没有绑定监听器或者后端云服务器没有关联ECS等原因，后端云服务器健康检查结果未知。
	OperatingStatus string `json:"operating_status"`

	Reason *MemberHealthCheckFailedReason `json:"reason,omitempty"`

	// **参数解释**：创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  **取值范围**：不涉及
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**：更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  **取值范围**：不涉及
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o MemberStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberStatus struct{}"
	}

	return strings.Join([]string{"MemberStatus", string(data)}, " ")
}
