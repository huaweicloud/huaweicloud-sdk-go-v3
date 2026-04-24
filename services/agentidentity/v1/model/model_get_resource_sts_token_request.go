package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceStsTokenRequest Request Object
type GetResourceStsTokenRequest struct {
	Body *GetResourceStsTokenRequestBody `json:"body,omitempty"`
}

func (o GetResourceStsTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceStsTokenRequest struct{}"
	}

	return strings.Join([]string{"GetResourceStsTokenRequest", string(data)}, " ")
}
