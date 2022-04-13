package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRecordSetRequest struct {
	ZoneId string `json:"zone_id"`

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
