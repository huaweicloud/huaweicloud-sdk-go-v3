package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetUserPasswordRequest Request Object
type ResetUserPasswordRequest struct {
}

func (o ResetUserPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetUserPasswordRequest", string(data)}, " ")
}
