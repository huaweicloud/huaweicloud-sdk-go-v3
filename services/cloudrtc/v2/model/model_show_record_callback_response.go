package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRecordCallbackResponse struct {

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	RecordCallback *AppCallbackUrl `json:"record_callback,omitempty" xml:"record_callback"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRecordCallbackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordCallbackResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordCallbackResponse", string(data)}, " ")
}
