package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecordSetsRequest Request Object
type DeleteRecordSetsRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	// 记录集ID。
	RecordsetId string `json:"recordset_id"`
}

func (o DeleteRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetsRequest", string(data)}, " ")
}
