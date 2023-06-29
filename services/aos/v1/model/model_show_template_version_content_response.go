package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateVersionContentResponse Response Object
type ShowTemplateVersionContentResponse struct {
	Location       *string `json:"Location,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTemplateVersionContentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateVersionContentResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateVersionContentResponse", string(data)}, " ")
}
