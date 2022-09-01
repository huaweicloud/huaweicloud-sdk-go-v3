package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAutoRecordResponse struct {

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	AutoRecordMode *AppAutoRecordMode `json:"auto_record_mode,omitempty" xml:"auto_record_mode"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAutoRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoRecordResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutoRecordResponse", string(data)}, " ")
}
