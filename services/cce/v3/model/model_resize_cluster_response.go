package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterResponse Response Object
type ResizeClusterResponse struct {

	// **参数解释**： 任务ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	JobID *string `json:"jobID,omitempty"`

	// **参数解释**： 包周期集群变更规格订单ID **约束限制**： v2接口不支持包周期集群规格变更 **取值范围**： 不涉及 **默认取值**： 不涉及
	OrderID        *string `json:"orderID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterResponse struct{}"
	}

	return strings.Join([]string{"ResizeClusterResponse", string(data)}, " ")
}
