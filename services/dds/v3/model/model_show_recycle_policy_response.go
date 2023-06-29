package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecyclePolicyResponse Response Object
type ShowRecyclePolicyResponse struct {
	RecyclePolicy  *RecyclePolicy `json:"recycle_policy,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowRecyclePolicyResponse", string(data)}, " ")
}
