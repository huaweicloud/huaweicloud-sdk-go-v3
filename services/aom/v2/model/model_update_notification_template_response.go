package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationTemplateResponse Response Object
type UpdateNotificationTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationTemplateResponse", string(data)}, " ")
}
