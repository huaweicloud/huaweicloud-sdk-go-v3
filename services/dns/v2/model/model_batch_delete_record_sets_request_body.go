package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRecordSetsRequestBody struct {

	// 域名的类型，取值为public或private。
	ZoneType string `json:"zone_type"`

	// 待删除的记录集ID列表。 最多支持100个。
	RecordsetIds []string `json:"recordset_ids"`
}

func (o BatchDeleteRecordSetsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetsRequestBody", string(data)}, " ")
}
