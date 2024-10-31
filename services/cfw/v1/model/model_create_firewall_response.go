package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFirewallResponse Response Object
type CreateFirewallResponse struct {

	// 实例创建的任务id。仅创建按需实例时会返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 订单号，创建包年包月时返回该参数。
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
