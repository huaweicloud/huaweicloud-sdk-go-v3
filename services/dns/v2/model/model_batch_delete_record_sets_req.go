package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRecordSetsReq struct {

	// Zone的类型，取值为public或private。
	ZoneType *string `json:"zone_type,omitempty"`

	// 待删除的Record Set ID列表。 最多支持100个。
	RecordsetIds *[]string `json:"recordset_ids,omitempty"`
}

func (o BatchDeleteRecordSetsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetsReq", string(data)}, " ")
}
