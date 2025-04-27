package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordSetRequest Request Object
type ShowRecordSetRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	// 记录集ID。
	RecordsetId string `json:"recordset_id"`
}

func (o ShowRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetRequest", string(data)}, " ")
}
