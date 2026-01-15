package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDbEncryptSecurityGroupResponse Response Object
type ChangeDbEncryptSecurityGroupResponse struct {

	// - 0：正常 - 非0：异常
	Code *int32 `json:"code,omitempty"`

	// 备注
	Description *string `json:"description,omitempty"`

	// 订单号
	OrderId *string `json:"order_id,omitempty"`

	// 任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeDbEncryptSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDbEncryptSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"ChangeDbEncryptSecurityGroupResponse", string(data)}, " ")
}
