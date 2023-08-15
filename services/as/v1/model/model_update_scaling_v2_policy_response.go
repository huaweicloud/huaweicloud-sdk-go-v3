package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScalingV2PolicyResponse Response Object
type UpdateScalingV2PolicyResponse struct {

	// 伸缩策略ID。
	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o UpdateScalingV2PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingV2PolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateScalingV2PolicyResponse", string(data)}, " ")
}
