package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRecordSetWithLineRequest struct {

	// 所属zone的ID。
	ZoneId string `json:"zone_id"`

	RecordsetId string `json:"recordset_id"`
}

func (o ShowRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetWithLineRequest", string(data)}, " ")
}
