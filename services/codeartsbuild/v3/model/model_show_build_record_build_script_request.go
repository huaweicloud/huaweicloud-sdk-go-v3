package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildRecordBuildScriptRequest Request Object
type ShowBuildRecordBuildScriptRequest struct {

	// 记录ID,36位数字、小写字母、'-'组组合。
	RecordId string `json:"record_id"`
}

func (o ShowBuildRecordBuildScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildRecordBuildScriptRequest struct{}"
	}

	return strings.Join([]string{"ShowBuildRecordBuildScriptRequest", string(data)}, " ")
}
