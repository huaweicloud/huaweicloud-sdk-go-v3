package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScreenRecordsRequest Request Object
type UpdateScreenRecordsRequest struct {

	// 录屏记录UUID。
	RecordId string `json:"record_id"`

	Body *UpdateScreenRecordsRequestBody `json:"body,omitempty"`
}

func (o UpdateScreenRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScreenRecordsRequest struct{}"
	}

	return strings.Join([]string{"UpdateScreenRecordsRequest", string(data)}, " ")
}
