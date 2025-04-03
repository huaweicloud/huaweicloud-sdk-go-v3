package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDeviceAuthorizationRequest Request Object
type StartDeviceAuthorizationRequest struct {
	Body *StartDeviceAuthorizationReqBody `json:"body,omitempty"`
}

func (o StartDeviceAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDeviceAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"StartDeviceAuthorizationRequest", string(data)}, " ")
}
