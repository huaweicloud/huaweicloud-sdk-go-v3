package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendNotificationsRequest Request Object
type SendNotificationsRequest struct {
	Body *SendNotificationsReq `json:"body,omitempty"`
}

func (o SendNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendNotificationsRequest struct{}"
	}

	return strings.Join([]string{"SendNotificationsRequest", string(data)}, " ")
}
