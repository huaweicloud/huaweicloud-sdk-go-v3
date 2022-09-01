package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteRecordSetWithLineRequest struct {

	// 所属zone的ID。
	ZoneId string `json:"zone_id" xml:"zone_id"`

	Body *BatchDeleteRSetWithLineReq `json:"body,omitempty" xml:"body"`
}

func (o BatchDeleteRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetWithLineRequest", string(data)}, " ")
}
