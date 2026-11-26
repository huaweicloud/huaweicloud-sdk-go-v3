package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApproverListApprovers struct {

	// 审核人ID
	Id *string `json:"id,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// 部门
	Department *string `json:"department,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 角色
	Character *string `json:"character,omitempty"`

	// 审核人名字
	ApproverName *string `json:"approver_name,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 电话号码
	PhoneNumber *string `json:"phone_number,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间
	CreateTime *int32 `json:"create_time,omitempty"`

	// APP名称
	AppName *string `json:"app_name,omitempty"`

	// 主题URN
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 审核人类型
	ApproverType *int32 `json:"approver_type,omitempty"`
}

func (o ApproverListApprovers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproverListApprovers struct{}"
	}

	return strings.Join([]string{"ApproverListApprovers", string(data)}, " ")
}
