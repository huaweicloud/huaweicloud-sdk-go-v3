package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqWorkItemDto **参数解释：** 工作项数据。
type ReqWorkItemDto struct {

	// **参数解释：** 工作项ID。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 工作项标题。
	Subject *string `json:"subject,omitempty"`

	// **参数解释：** 工作项编号， scrum类型项目该值为空。
	Sequence *string `json:"sequence,omitempty"`

	Tracker *ReqWorkItemBasicDto `json:"tracker,omitempty"`

	Status *ReqWorkItemBasicDto `json:"status,omitempty"`

	Priority *ReqWorkItemBasicDto `json:"priority,omitempty"`
}

func (o ReqWorkItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqWorkItemDto struct{}"
	}

	return strings.Join([]string{"ReqWorkItemDto", string(data)}, " ")
}
