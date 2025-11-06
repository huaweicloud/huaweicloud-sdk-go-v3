package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotificationTemplateResponse Response Object
type CreateNotificationTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateNotificationTemplateResponse", string(data)}, " ")
}
