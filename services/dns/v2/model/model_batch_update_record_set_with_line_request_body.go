package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateRecordSetWithLineRequestBody struct {

	// RecordSet 列表。
	Recordsets []BatchUpdateRecordSet `json:"recordsets"`
}

func (o BatchUpdateRecordSetWithLineRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRecordSetWithLineRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateRecordSetWithLineRequestBody", string(data)}, " ")
}
