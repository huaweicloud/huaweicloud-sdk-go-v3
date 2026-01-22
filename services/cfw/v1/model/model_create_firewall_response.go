package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFirewallResponse Response Object
type CreateFirewallResponse struct {

	// **参数解释**： 实例创建的任务id **取值范围**： 不涉及
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 订单号 **取值范围**： 不涉及
	OrderId *string `json:"order_id,omitempty"`

	Data           *CreateFirewallReq `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallResponse struct{}"
	}

	return strings.Join([]string{"CreateFirewallResponse", string(data)}, " ")
}
