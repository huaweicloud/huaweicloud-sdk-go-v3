package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPtrRecordSetRequest struct {
	// 租户的区域信息。

	Region string `json:"region"`
	// 弹性IP的ID。

	FloatingipId string `json:"floatingip_id"`
}

func (o ShowPtrRecordSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPtrRecordSetRequest struct{}"
	}

	return strings.Join([]string{"ShowPtrRecordSetRequest", string(data)}, " ")
}
