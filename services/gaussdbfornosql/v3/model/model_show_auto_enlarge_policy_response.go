package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoEnlargePolicyResponse Response Object
type ShowAutoEnlargePolicyResponse struct {
	Policy         *DiskAutoExpansionPolicy `json:"policy,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowAutoEnlargePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoEnlargePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoEnlargePolicyResponse", string(data)}, " ")
}
