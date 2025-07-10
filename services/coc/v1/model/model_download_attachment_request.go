package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAttachmentRequest Request Object
type DownloadAttachmentRequest struct {

	// **参数解释：** 工单类型，此处传递固定值incident。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TicketType string `json:"ticket_type"`

	// **参数解释：** 需要下载的文件唯一ID。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DocId string `json:"doc_id"`
}

func (o DownloadAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DownloadAttachmentRequest", string(data)}, " ")
}
