package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqWorkItemBasicDto **参数解释：** 工作项基础对象，用于记录工作项的状态、类型等属性的ID和名称。
type ReqWorkItemBasicDto struct {

	// **参数解释：** 属性ID。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 属性名称。
	Name *string `json:"name,omitempty"`
}

func (o ReqWorkItemBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqWorkItemBasicDto struct{}"
	}

	return strings.Join([]string{"ReqWorkItemBasicDto", string(data)}, " ")
}
