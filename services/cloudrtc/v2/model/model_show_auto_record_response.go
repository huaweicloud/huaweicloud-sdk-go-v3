package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoRecordResponse Response Object
type ShowAutoRecordResponse struct {

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	AutoRecordMode *AppAutoRecordMode `json:"auto_record_mode,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAutoRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoRecordResponse", string(data)}, " ")
}
