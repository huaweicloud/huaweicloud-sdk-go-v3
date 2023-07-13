package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotificationTemplateResponse Response Object
type DeleteNotificationTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteNotificationTemplateResponse", string(data)}, " ")
}
