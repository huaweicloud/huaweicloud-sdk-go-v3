package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteScreenRecordsRequestBody struct {

	// 待删除的录屏ID列表。
	ScreenRecords *[]BatchDeleteScreenRecordsRequestBodyScreenRecords `json:"screen_records,omitempty"`
}

func (o BatchDeleteScreenRecordsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScreenRecordsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteScreenRecordsRequestBody", string(data)}, " ")
}
