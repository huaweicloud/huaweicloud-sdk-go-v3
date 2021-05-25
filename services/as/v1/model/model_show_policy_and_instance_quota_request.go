package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPolicyAndInstanceQuotaRequest struct {
	// 伸缩组ID。

	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ShowPolicyAndInstanceQuotaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPolicyAndInstanceQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyAndInstanceQuotaRequest", string(data)}, " ")
}
