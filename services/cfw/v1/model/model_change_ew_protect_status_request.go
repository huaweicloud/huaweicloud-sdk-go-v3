package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeEwProtectStatusRequest struct {
	Body *ChangeProtectStatusRequestBody `json:"body,omitempty"`
}

func (o ChangeEwProtectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEwProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeEwProtectStatusRequest", string(data)}, " ")
}
