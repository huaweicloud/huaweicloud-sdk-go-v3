package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildRecordRequest Request Object
type ShowBuildRecordRequest struct {

	// 记录ID,36位数字、小写字母、'-'组组合。
	RecordId string `json:"record_id"`
}

func (o ShowBuildRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowBuildRecordRequest", string(data)}, " ")
}
