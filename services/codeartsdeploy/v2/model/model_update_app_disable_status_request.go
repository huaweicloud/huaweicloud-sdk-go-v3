package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppDisableStatusRequest Request Object
type UpdateAppDisableStatusRequest struct {

	// 应用id
	AppId string `json:"app_id"`

	Body *UpdateAppDisableStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateAppDisableStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppDisableStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppDisableStatusRequest", string(data)}, " ")
}
