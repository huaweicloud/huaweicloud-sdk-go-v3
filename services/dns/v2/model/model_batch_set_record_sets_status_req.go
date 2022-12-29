package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetRecordSetsStatusReq struct {

	// 待设置Record Se状态，当前仅支持DISABLE或ENABLE。
	Status *string `json:"status,omitempty"`

	// 待设置Record Set ID列表。 最多支持50个。
	RecordsetIds *[]string `json:"recordset_ids,omitempty"`
}

func (o BatchSetRecordSetsStatusReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetRecordSetsStatusReq struct{}"
	}

	return strings.Join([]string{"BatchSetRecordSetsStatusReq", string(data)}, " ")
}
