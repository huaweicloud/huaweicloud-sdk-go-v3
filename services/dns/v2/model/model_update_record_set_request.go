package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecordSetRequest Request Object
type UpdateRecordSetRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	// 记录集ID。
	RecordsetId string `json:"recordset_id"`

	Body *UpdateRecordSetReq `json:"body,omitempty"`
}

func (o UpdateRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetRequest", string(data)}, " ")
}
