package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFullSpeedRecordConfigRequest Request Object
type UpdateFullSpeedRecordConfigRequest struct {

	// 录屏记录UUID。
	RecordId string `json:"record_id"`

	Body *UpdateScreenRecordsRequestBody `json:"body,omitempty"`
}

func (o UpdateFullSpeedRecordConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFullSpeedRecordConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateFullSpeedRecordConfigRequest", string(data)}, " ")
}
