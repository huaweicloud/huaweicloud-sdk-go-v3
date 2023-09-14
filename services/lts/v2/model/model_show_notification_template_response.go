package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNotificationTemplateResponse Response Object
type ShowNotificationTemplateResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowNotificationTemplateResponse", string(data)}, " ")
}
