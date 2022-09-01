package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetRecordSetsStatusRequest struct {
	RecordsetId string `json:"recordset_id" xml:"recordset_id"`

	Body *SetRecordSetsStatusReq `json:"body,omitempty" xml:"body"`
}

func (o SetRecordSetsStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecordSetsStatusRequest struct{}"
	}

	return strings.Join([]string{"SetRecordSetsStatusRequest", string(data)}, " ")
}
