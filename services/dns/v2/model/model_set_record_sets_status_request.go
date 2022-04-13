package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetRecordSetsStatusRequest struct {
	RecordsetId string `json:"recordset_id"`

	Body *SetRecordSetsStatusReq `json:"body,omitempty"`
}

func (o SetRecordSetsStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecordSetsStatusRequest struct{}"
	}

	return strings.Join([]string{"SetRecordSetsStatusRequest", string(data)}, " ")
}
