package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyEngineAttachmentsResponse Response Object
type ListPolicyEngineAttachmentsResponse struct {
	Attachments *[]PolicyEngineAttachmentSummary `json:"attachments,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPolicyEngineAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyEngineAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyEngineAttachmentsResponse", string(data)}, " ")
}
