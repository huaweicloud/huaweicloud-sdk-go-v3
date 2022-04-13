package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHostGroupRequest struct {
	Body *UpdateHostGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateHostGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostGroupRequest", string(data)}, " ")
}
