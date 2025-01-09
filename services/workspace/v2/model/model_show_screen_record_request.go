package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScreenRecordRequest Request Object
type ShowScreenRecordRequest struct {

	// 录屏记录UUID。
	RecordId string `json:"record_id"`
}

func (o ShowScreenRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScreenRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowScreenRecordRequest", string(data)}, " ")
}
