package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationMasksResponse Response Object
type UpdateNotificationMasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNotificationMasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationMasksResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationMasksResponse", string(data)}, " ")
}
