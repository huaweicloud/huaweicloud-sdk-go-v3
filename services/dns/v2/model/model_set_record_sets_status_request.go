package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRecordSetsStatusRequest Request Object
type SetRecordSetsStatusRequest struct {

	// 记录集ID。
	RecordsetId string `json:"recordset_id"`

	Body *SetRecordSetsStatusRequestBody `json:"body,omitempty"`
}

func (o SetRecordSetsStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecordSetsStatusRequest struct{}"
	}

	return strings.Join([]string{"SetRecordSetsStatusRequest", string(data)}, " ")
}
