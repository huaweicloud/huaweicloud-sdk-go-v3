package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 事件内容
type OrderAlertIncidentContent struct {

	// 事件名称
	Title *string `json:"title,omitempty"`

	// 事件类型
	TypeCategory *string `json:"type_category,omitempty"`

	// 证据列表
	EvidenceList *[]string `json:"evidence_list,omitempty"`

	// 评论列表
	NoteList *[]string `json:"note_list,omitempty"`

	// 附件列表
	AttachmentList *[]string `json:"attachment_list,omitempty"`

	IncidentType *OrderAlertIncidentContentIncidentType `json:"incident_type,omitempty"`

	// Id value
	Description *string `json:"description,omitempty"`
}

func (o OrderAlertIncidentContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderAlertIncidentContent struct{}"
	}

	return strings.Join([]string{"OrderAlertIncidentContent", string(data)}, " ")
}
