package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildRecordFullStagesRequest Request Object
type ShowBuildRecordFullStagesRequest struct {

	// 记录ID,36位数字、小写字母、'-'组组合。
	RecordId string `json:"record_id"`

	// 是否联级获取steps
	Cascade *bool `json:"cascade,omitempty"`
}

func (o ShowBuildRecordFullStagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildRecordFullStagesRequest struct{}"
	}

	return strings.Join([]string{"ShowBuildRecordFullStagesRequest", string(data)}, " ")
}
