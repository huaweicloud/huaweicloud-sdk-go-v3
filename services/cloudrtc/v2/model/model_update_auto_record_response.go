package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAutoRecordResponse struct {
	// 应用id

	AppId *string `json:"app_id,omitempty"`

	AutoRecordMode *AppAutoRecordMode `json:"auto_record_mode,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAutoRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoRecordResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutoRecordResponse", string(data)}, " ")
}
