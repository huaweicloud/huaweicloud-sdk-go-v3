package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTagForProtectedIpRequest Request Object
type UpdateTagForProtectedIpRequest struct {
	Body *UpdateProtectedIpBody `json:"body,omitempty"`
}

func (o UpdateTagForProtectedIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTagForProtectedIpRequest struct{}"
	}

	return strings.Join([]string{"UpdateTagForProtectedIpRequest", string(data)}, " ")
}
