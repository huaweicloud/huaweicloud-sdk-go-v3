package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionListResp 订阅任务列表响应体。
type SubscriptionListResp struct {

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 任务状态。
	Status string `json:"status"`

	// 任务创建时间。
	CreateTime string `json:"create_time"`

	// 消费开始时间。
	BeginTime string `json:"begin_time"`

	// 任务描述。
	Description string `json:"description"`

	// 当前时间。
	NowTime string `json:"now_time"`

	JobAction *JobActions `json:"job_action"`

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o SubscriptionListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionListResp struct{}"
	}

	return strings.Join([]string{"SubscriptionListResp", string(data)}, " ")
}
