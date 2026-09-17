package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePlanStatusRequest Request Object
type ChangePlanStatusRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 发布/迭代唯一ID
	PlanId string `json:"plan_id"`

	Body *CompleteSprintVo `json:"body,omitempty"`
}

func (o ChangePlanStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePlanStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangePlanStatusRequest", string(data)}, " ")
}
