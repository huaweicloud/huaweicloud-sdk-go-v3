package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRecordSetWithLineRequest struct {
	ZoneId string `json:"zone_id"`

	Body *CreateRecordSetWithLineReq `json:"body,omitempty"`
}

func (o CreateRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithLineRequest", string(data)}, " ")
}
