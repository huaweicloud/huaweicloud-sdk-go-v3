package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyProviderRequest Request Object
type VerifyProviderRequest struct {
	Body *VerifyProviderReq `json:"body,omitempty"`
}

func (o VerifyProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyProviderRequest struct{}"
	}

	return strings.Join([]string{"VerifyProviderRequest", string(data)}, " ")
}
