package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecipientsStatusResponse Response Object
type ListRecipientsStatusResponse struct {
	Body *[]EmailStatusInfo `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRecipientsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecipientsStatusResponse struct{}"
	}

	return strings.Join([]string{"ListRecipientsStatusResponse", string(data)}, " ")
}
