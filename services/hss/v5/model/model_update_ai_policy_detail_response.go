package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiPolicyDetailResponse Response Object
type UpdateAiPolicyDetailResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAiPolicyDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiPolicyDetailResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiPolicyDetailResponse", string(data)}, " ")
}
