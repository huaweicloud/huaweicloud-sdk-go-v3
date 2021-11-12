package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRecordSetByZoneRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o ShowRecordSetByZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetByZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetByZoneRequest", string(data)}, " ")
}
