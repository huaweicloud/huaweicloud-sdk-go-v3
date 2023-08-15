package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemStatusFlowVo 工作项流转数据对象
type WorkItemStatusFlowVo struct {

	//  父状态的名称
	ParentName *string `json:"parent_name,omitempty"`

	// 父状态的类型
	ParentType *string `json:"parent_type,omitempty"`

	// 状态id
	StatusId *string `json:"status_id,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`

	// 状态类型
	StatusType *string `json:"status_type,omitempty"`

	// 流转到的数据
	DirectTo *[]StatusFlowDirectToVo `json:"direct_to,omitempty"`

	// 处理人的uuid
	AssignTo *string `json:"assign_to,omitempty"`

	// 评论内容
	Comment *string `json:"comment,omitempty"`

	// 处理人是否必填
	RequiredAssign *bool `json:"required_assign,omitempty"`

	// 评论是否必填
	RequiredNotes *bool `json:"required_notes,omitempty"`

	// 是否是字段值，true 处理人的信息是字段值， false 处理人的值是用户的信息,固定值
	FieldType *bool `json:"field_type,omitempty"`

	// 父状态的uuid
	ParentId *string `json:"parent_id,omitempty"`
}

func (o WorkItemStatusFlowVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemStatusFlowVo struct{}"
	}

	return strings.Join([]string{"WorkItemStatusFlowVo", string(data)}, " ")
}
