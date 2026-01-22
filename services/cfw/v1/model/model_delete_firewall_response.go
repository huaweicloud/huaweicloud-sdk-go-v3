package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFirewallResponse Response Object
type DeleteFirewallResponse struct {

	// **参数解释**： 删除防火墙时生成的任务的job_id **约束限制**： 不涉及
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFirewallResponse struct{}"
	}

	return strings.Join([]string{"DeleteFirewallResponse", string(data)}, " ")
}
