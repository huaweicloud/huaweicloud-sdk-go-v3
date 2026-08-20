package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProcessInstanceReq struct {

	// 标题
	Title string `json:"title"`

	// 描述
	Description string `json:"description"`

	// 类别
	Category string `json:"category"`

	// 是否需要决策人审批
	NeedApproval *bool `json:"need_approval,omitempty"`

	// 计划完成日期时间戳，不可早于计划开始日期
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// 计划开始日期时间戳，不可晚于计划完成日期
	PlanStartDate *string `json:"plan_start_date,omitempty"`

	// 状态
	Status string `json:"status"`

	// 抄送人列表
	Cc *[]string `json:"cc,omitempty"`

	// 关联wiki
	AttachWikis *[]string `json:"attachWikis,omitempty"`

	// 关联文件
	AttachDocuments *[]string `json:"attachDocuments,omitempty"`

	// 决策人
	Ccbs *[]CreateProcessInstanceReqCcbs `json:"ccbs,omitempty"`

	// 评审专家
	Opinions *[]CreateProcessInstanceReqOpinions `json:"opinions,omitempty"`

	// 评审对象
	Cos *[]CreateProcessInstanceReqCos `json:"cos,omitempty"`

	// 关联文件名
	LocalAttachmentNames *[]string `json:"local_attachment_names,omitempty"`
}

func (o CreateProcessInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProcessInstanceReq struct{}"
	}

	return strings.Join([]string{"CreateProcessInstanceReq", string(data)}, " ")
}
