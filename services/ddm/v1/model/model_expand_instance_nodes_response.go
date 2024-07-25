package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandInstanceNodesResponse Response Object
type ExpandInstanceNodesResponse struct {

	// DDM实例ID。
	InstanceId *string `json:"instanceId,omitempty"`

	// DDM实例名称,仅按需实例时会返回该参数。
	InstanceName *string `json:"instanceName,omitempty"`

	// 任务ID,仅按需实例时会返回该参数。
	JobId *string `json:"jobId,omitempty"`

	// 订单号,仅包年包月实例时返回该参数。
	OrderId        *string `json:"orderId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandInstanceNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodesResponse struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodesResponse", string(data)}, " ")
}
