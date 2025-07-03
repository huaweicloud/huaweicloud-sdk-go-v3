package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAttachmentRequest Request Object
type CreateAttachmentRequest struct {

	// **参数解释：** 需要创建的工单类型，此处为固定值incident。此处的ticket_type传递的值为固定值incident，但是创建变更单或者问题单时，如需上传附件使用的依然是此接口。 **约束限制：** 不涉及 **取值范围：** 固定值incident **默认取值：** 不涉及
	TicketType string `json:"ticket_type"`

	Body *CreateAttachmentRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAttachmentRequest struct{}"
	}

	return strings.Join([]string{"CreateAttachmentRequest", string(data)}, " ")
}
