package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeObsRequest Request Object
type AuthorizeObsRequest struct {
	Body *AuthorizeObsReq `json:"body,omitempty"`
}

func (o AuthorizeObsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeObsRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeObsRequest", string(data)}, " ")
}
