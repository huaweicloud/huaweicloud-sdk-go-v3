package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeploymentsRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *Deployment `json:"body,omitempty"`
}

func (o CreateDeploymentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentsRequest struct{}"
	}

	return strings.Join([]string{"CreateDeploymentsRequest", string(data)}, " ")
}
