package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimSendReportsResponse Response Object
type ListAimSendReportsResponse struct {

	// 响应码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Data           *ListAimSendReportsMode `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListAimSendReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimSendReportsResponse struct{}"
	}

	return strings.Join([]string{"ListAimSendReportsResponse", string(data)}, " ")
}
