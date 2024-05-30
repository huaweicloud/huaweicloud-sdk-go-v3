package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalBatchParam struct {

	// 业务信息列表。
	BizInfos []BizInfoVo `json:"biz_infos"`

	// 审批人ID。
	ApproverUserId string `json:"approver_user_id"`

	// 审批人姓名。
	ApproverUserName string `json:"approver_user_name"`

	// 审批人邮箱，仅在创建审批人时填写。
	Email *string `json:"email,omitempty"`

	// 快速审批，非正式场景，用于快速上手体验，仅在当前用户有审批权限时提供。
	FastApproval *bool `json:"fast_approval,omitempty"`

	// 作业调度时间。格式参照：30_18，表示18点30分。
	ScheduleTime *string `json:"schedule_time,omitempty"`

	EnvType *EnvTypeEnum `json:"env_type,omitempty"`
}

func (o ApprovalBatchParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalBatchParam struct{}"
	}

	return strings.Join([]string{"ApprovalBatchParam", string(data)}, " ")
}
