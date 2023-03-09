package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTemplateVersionContentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowTemplateVersionContentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateVersionContentResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateVersionContentResponse", string(data)}, " ")
}
