package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateRecordSetsTaskRequestBody struct {

	// 批量创建记录集列表。
	Recordsets []BatchCreateRecordSetsTaskItem `json:"recordsets"`
}

func (o BatchCreateRecordSetsTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateRecordSetsTaskRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateRecordSetsTaskRequestBody", string(data)}, " ")
}
