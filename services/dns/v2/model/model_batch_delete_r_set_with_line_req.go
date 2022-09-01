package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRSetWithLineReq struct {

	// Record Set ID列表。最多支持100个。
	RecordsetIds []string `json:"recordset_ids" xml:"recordset_ids"`
}

func (o BatchDeleteRSetWithLineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRSetWithLineReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteRSetWithLineReq", string(data)}, " ")
}
