package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTokenRequest Request Object
type CreateTokenRequest struct {
	Body *CreateTokenReqBody `json:"body,omitempty"`
}

func (o CreateTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateTokenRequest", string(data)}, " ")
}
