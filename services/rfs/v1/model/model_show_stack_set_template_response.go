package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStackSetTemplateResponse Response Object
type ShowStackSetTemplateResponse struct {
	Location       *string `json:"Location,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStackSetTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStackSetTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowStackSetTemplateResponse", string(data)}, " ")
}
