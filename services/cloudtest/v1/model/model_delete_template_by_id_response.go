package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateByIdResponse Response Object
type DeleteTemplateByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTemplateByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateByIdResponse", string(data)}, " ")
}
