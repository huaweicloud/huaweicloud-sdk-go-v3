package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSystemTemplateDetailResponse Response Object
type ShowSystemTemplateDetailResponse struct {
	ProvidedActions *[]ProvidedAction `json:"provided_actions,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSystemTemplateDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSystemTemplateDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSystemTemplateDetailResponse", string(data)}, " ")
}
