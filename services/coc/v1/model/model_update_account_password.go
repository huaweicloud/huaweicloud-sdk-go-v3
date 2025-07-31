package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccountPassword 回写改密结果的信息
type UpdateAccountPassword struct {

	// 改密计划唯一UUID
	PlanId string `json:"plan_id"`

	// 改密资源id
	ResourceId string `json:"resource_id"`

	// 改密账号名称
	AccountName string `json:"account_name"`

	// 改密结果状态码
	StatusCode string `json:"status_code"`

	// 改密结果详情
	Detail string `json:"detail"`
}

func (o UpdateAccountPassword) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccountPassword struct{}"
	}

	return strings.Join([]string{"UpdateAccountPassword", string(data)}, " ")
}
