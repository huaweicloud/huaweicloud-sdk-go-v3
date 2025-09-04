package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetUserProfileRequest Request Object
type ResetUserProfileRequest struct {
	Body *ResetUserProfileReq `json:"body,omitempty"`
}

func (o ResetUserProfileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserProfileRequest struct{}"
	}

	return strings.Join([]string{"ResetUserProfileRequest", string(data)}, " ")
}
