package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TermAssignmentHeader 关联标签详情
type TermAssignmentHeader struct {

	// 信任id
	Confidence *int32 `json:"confidence,omitempty"`

	// 管理员
	Steward *string `json:"steward,omitempty"`

	// 来源
	Source *string `json:"source,omitempty"`

	// 状态 枚举值：DISCOVERED、PROPOSED、IMPORTED、VALIDATED、DEPRECATED、OBSOLETE、OTHER
	Status *string `json:"status,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 表达式
	Expression *string `json:"expression,omitempty"`

	// 展示信息
	DisplayText *string `json:"display_text,omitempty"`

	// 标签guid
	TermGuid *string `json:"term_guid,omitempty"`

	// 关联guid
	RelationGuid *string `json:"relation_guid,omitempty"`
}

func (o TermAssignmentHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TermAssignmentHeader struct{}"
	}

	return strings.Join([]string{"TermAssignmentHeader", string(data)}, " ")
}
