package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordCallbackResponse Response Object
type ShowRecordCallbackResponse struct {

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	RecordCallback *AppCallbackUrl `json:"record_callback,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRecordCallbackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordCallbackResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordCallbackResponse", string(data)}, " ")
}
