package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadLogByRecordIdRequest Request Object
type DownloadLogByRecordIdRequest struct {

	// 记录ID,36位数字、小写字母、'-'组组合。
	RecordId string `json:"record_id"`
}

func (o DownloadLogByRecordIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadLogByRecordIdRequest struct{}"
	}

	return strings.Join([]string{"DownloadLogByRecordIdRequest", string(data)}, " ")
}
