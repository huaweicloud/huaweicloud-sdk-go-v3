package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMessageNotificationPolicyResponse Response Object
type DeleteMessageNotificationPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMessageNotificationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMessageNotificationPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteMessageNotificationPolicyResponse", string(data)}, " ")
}
