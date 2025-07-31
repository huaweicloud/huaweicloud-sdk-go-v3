package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountChangePwdPlan 改密计划的信息
type AccountChangePwdPlan struct {

	// 改密计划唯一UUID
	PlanId string `json:"plan_id"`

	// 改密资源id
	ResourceId string `json:"resource_id"`

	// 改密账号名称
	AccountName string `json:"account_name"`

	// 需要修改的密码
	Password string `json:"password"`

	// 资源所属的项目id
	ProjectId string `json:"project_id"`

	// 改密计划创建时间
	CreateTime int64 `json:"create_time"`
}

func (o AccountChangePwdPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountChangePwdPlan struct{}"
	}

	return strings.Join([]string{"AccountChangePwdPlan", string(data)}, " ")
}
