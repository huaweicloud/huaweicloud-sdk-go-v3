package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotificationSettingRequest Request Object
type CreateNotificationSettingRequest struct {
	Body *CreateNotificationSettingReqBody `json:"body,omitempty"`
}

func (o CreateNotificationSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationSettingRequest struct{}"
	}

	return strings.Join([]string{"CreateNotificationSettingRequest", string(data)}, " ")
}
