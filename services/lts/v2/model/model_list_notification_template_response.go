package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationTemplateResponse Response Object
type ListNotificationTemplateResponse struct {

	// 为一个html文本，需要进行相应的解析后展示
	Template       *string `json:"template,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationTemplateResponse", string(data)}, " ")
}
