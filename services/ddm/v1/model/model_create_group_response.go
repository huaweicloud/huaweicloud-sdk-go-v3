package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupResponse Response Object
type CreateGroupResponse struct {

	// DDM实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务ID，仅创建按需实例时会返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 订单号，创建包年包月时返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupResponse", string(data)}, " ")
}
