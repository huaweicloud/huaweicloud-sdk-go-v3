package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicTemplateInfoResponse Response Object
type ShowPublicTemplateInfoResponse struct {
	ProvidedActions *ProvidedAction `json:"provided_actions,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPublicTemplateInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicTemplateInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicTemplateInfoResponse", string(data)}, " ")
}
