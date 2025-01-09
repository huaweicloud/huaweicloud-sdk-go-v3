package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScriptRecordDetailRequest Request Object
type ShowScriptRecordDetailRequest struct {

	// 脚本执行记录ID。
	RecordId string `json:"record_id"`
}

func (o ShowScriptRecordDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptRecordDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowScriptRecordDetailRequest", string(data)}, " ")
}
