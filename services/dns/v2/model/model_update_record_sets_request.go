package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRecordSetsRequest struct {
	ZoneId string `json:"zone_id" xml:"zone_id"`

	RecordsetId string `json:"recordset_id" xml:"recordset_id"`

	Body *UpdateRecordSetsReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetsRequest", string(data)}, " ")
}
