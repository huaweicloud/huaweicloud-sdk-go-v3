package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateByIdResponse Response Object
type ShowTemplateByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTemplateByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateByIdResponse", string(data)}, " ")
}
