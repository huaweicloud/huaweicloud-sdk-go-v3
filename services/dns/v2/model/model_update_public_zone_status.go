package model

import (
	"encoding/json"

	"strings"
)

type UpdatePublicZoneStatus struct {
	// Zone状态。  取值范围：  ENABLE：启用解析 DISABLE：暂停解析

	Status string `json:"status"`
}

func (o UpdatePublicZoneStatus) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneStatus struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneStatus", string(data)}, " ")
}
