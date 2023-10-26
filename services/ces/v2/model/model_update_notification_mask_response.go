package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationMaskResponse Response Object
type UpdateNotificationMaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNotificationMaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationMaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationMaskResponse", string(data)}, " ")
}
