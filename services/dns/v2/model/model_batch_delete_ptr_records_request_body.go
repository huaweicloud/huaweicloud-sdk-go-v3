package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeletePtrRecordsRequestBody struct {

	// 待删除反向解析记录ID列表。 最多支持50个。
	FloatingipIds *[]string `json:"floatingip_ids,omitempty"`
}

func (o BatchDeletePtrRecordsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePtrRecordsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeletePtrRecordsRequestBody", string(data)}, " ")
}
