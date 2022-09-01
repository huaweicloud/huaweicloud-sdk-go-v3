package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteRecordSetsRequest struct {
	ZoneId string `json:"zone_id" xml:"zone_id"`

	RecordsetId string `json:"recordset_id" xml:"recordset_id"`
}

func (o DeleteRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetsRequest", string(data)}, " ")
}
