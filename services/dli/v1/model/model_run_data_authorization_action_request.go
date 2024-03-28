package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDataAuthorizationActionRequest Request Object
type RunDataAuthorizationActionRequest struct {
	Body *RunDataAuthorizationActionRequestBody `json:"body,omitempty"`
}

func (o RunDataAuthorizationActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDataAuthorizationActionRequest struct{}"
	}

	return strings.Join([]string{"RunDataAuthorizationActionRequest", string(data)}, " ")
}
