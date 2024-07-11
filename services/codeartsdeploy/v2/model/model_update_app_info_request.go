package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppInfoRequest Request Object
type UpdateAppInfoRequest struct {
	Body *UpdateAppInfoRequestBody `json:"body,omitempty"`
}

func (o UpdateAppInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppInfoRequest", string(data)}, " ")
}
