package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项流转数据对象
type WorkItemStatusFlowVo struct {

	//  父状态的名称
	ParentName *string `json:"parent_name,omitempty" xml:"parent_name"`

	// 父状态的类型
	ParentType *string `json:"parent_type,omitempty" xml:"parent_type"`

	// 状态id
	StatusId *string `json:"status_id,omitempty" xml:"status_id"`

	// 状态名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 状态类型
	StatusType *string `json:"status_type,omitempty" xml:"status_type"`

	// 流转到的数据
	DirectTo *[]StatusFlowDirectToVo `json:"direct_to,omitempty" xml:"direct_to"`

	// 处理人的uuid
	AssignTo *string `json:"assign_to,omitempty" xml:"assign_to"`

	// 评论内容
	Comment *string `json:"comment,omitempty" xml:"comment"`

	// 处理人是否必填
	RequiredAssign *bool `json:"required_assign,omitempty" xml:"required_assign"`

	// 评论是否必填
	RequiredNotes *bool `json:"required_notes,omitempty" xml:"required_notes"`

	// 是否是字段值，true 处理人的信息是字段值， false 处理人的值是用户的信息,固定值
	FieldType *bool `json:"field_type,omitempty" xml:"field_type"`

	// 父状态的uuid
	ParentId *string `json:"parent_id,omitempty" xml:"parent_id"`
}

func (o WorkItemStatusFlowVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemStatusFlowVo struct{}"
	}

	return strings.Join([]string{"WorkItemStatusFlowVo", string(data)}, " ")
}
