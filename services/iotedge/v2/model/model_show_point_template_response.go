package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPointTemplateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPointTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPointTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowPointTemplateResponse", string(data)}, " ")
}
