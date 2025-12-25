package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecipientsEmailStatusResponse Response Object
type UpdateRecipientsEmailStatusResponse struct {

	// 响应示例
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRecipientsEmailStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecipientsEmailStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecipientsEmailStatusResponse", string(data)}, " ")
}
