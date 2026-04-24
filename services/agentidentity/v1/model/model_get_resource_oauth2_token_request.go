package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceOauth2TokenRequest Request Object
type GetResourceOauth2TokenRequest struct {
	Body *GetResourceOauth2TokenRequestBody `json:"body,omitempty"`
}

func (o GetResourceOauth2TokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceOauth2TokenRequest struct{}"
	}

	return strings.Join([]string{"GetResourceOauth2TokenRequest", string(data)}, " ")
}
