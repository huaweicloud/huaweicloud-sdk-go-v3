package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthorizationRequest Request Object
type ShowAuthorizationRequest struct {
}

func (o ShowAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"ShowAuthorizationRequest", string(data)}, " ")
}
