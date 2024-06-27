package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateByPageResponse Response Object
type ShowTemplateByPageResponse struct {
	Code *string `json:"code,omitempty"`

	Data *BasePageInfoTemplateV2 `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTemplateByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateByPageResponse", string(data)}, " ")
}
