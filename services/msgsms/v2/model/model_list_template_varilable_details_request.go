package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplateVarilableDetailsRequest Request Object
type ListTemplateVarilableDetailsRequest struct {

	// 模板ID
	Id string `json:"id"`
}

func (o ListTemplateVarilableDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateVarilableDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateVarilableDetailsRequest", string(data)}, " ")
}
