package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeResourceRequest Request Object
type AuthorizeResourceRequest struct {
	Body *AuthorizeResourceRequestBody `json:"body,omitempty"`
}

func (o AuthorizeResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeResourceRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeResourceRequest", string(data)}, " ")
}
