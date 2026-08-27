package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProviderRequest Request Object
type CreateProviderRequest struct {
	Body *CreateProviderReq `json:"body,omitempty"`
}

func (o CreateProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProviderRequest struct{}"
	}

	return strings.Join([]string{"CreateProviderRequest", string(data)}, " ")
}
