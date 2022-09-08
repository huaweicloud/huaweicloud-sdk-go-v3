package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateRecordCallbackResponse struct {

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	RecordCallback *AppCallbackUrl `json:"record_callback,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRecordCallbackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordCallbackResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecordCallbackResponse", string(data)}, " ")
}
