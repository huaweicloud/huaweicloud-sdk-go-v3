package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueueSrcAssignUpdateDto struct {

	// 队列属性(0:默认,1:实时队列,2:离线队列), 当前只有yarn队列涉及。
	QueueAttr *int32 `json:"queue_attr,omitempty"`

	// 当前空间分配资源附加的描述信息。
	Description *string `json:"description,omitempty"`
}

func (o QueueSrcAssignUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueSrcAssignUpdateDto struct{}"
	}

	return strings.Join([]string{"QueueSrcAssignUpdateDto", string(data)}, " ")
}
