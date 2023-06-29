package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowThirdTemplateInfoResponse Response Object
type ShowThirdTemplateInfoResponse struct {
	ProvidedActions *ProvidedAction `json:"provided_actions,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowThirdTemplateInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowThirdTemplateInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowThirdTemplateInfoResponse", string(data)}, " ")
}
