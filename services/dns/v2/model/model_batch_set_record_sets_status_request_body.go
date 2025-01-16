package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetRecordSetsStatusRequestBody struct {

	// 待设置Record Set状态，当前仅支持DISABLE或ENABLE。
	Status string `json:"status"`

	// 待设置Record Set ID列表。 最多支持50个。
	RecordsetIds []string `json:"recordset_ids"`
}

func (o BatchSetRecordSetsStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetRecordSetsStatusRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSetRecordSetsStatusRequestBody", string(data)}, " ")
}
