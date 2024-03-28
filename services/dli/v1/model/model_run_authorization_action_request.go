package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunAuthorizationActionRequest Request Object
type RunAuthorizationActionRequest struct {
	Body *RunAuthorizationActionRequestBody `json:"body,omitempty"`
}

func (o RunAuthorizationActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAuthorizationActionRequest struct{}"
	}

	return strings.Join([]string{"RunAuthorizationActionRequest", string(data)}, " ")
}
