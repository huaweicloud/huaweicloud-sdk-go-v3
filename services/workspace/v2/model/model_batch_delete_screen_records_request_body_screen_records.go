package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteScreenRecordsRequestBodyScreenRecords struct {

	// 主键UUID。
	Id *string `json:"id,omitempty"`
}

func (o BatchDeleteScreenRecordsRequestBodyScreenRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScreenRecordsRequestBodyScreenRecords struct{}"
	}

	return strings.Join([]string{"BatchDeleteScreenRecordsRequestBodyScreenRecords", string(data)}, " ")
}
