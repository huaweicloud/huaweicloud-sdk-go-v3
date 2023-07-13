package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMyActionTemplateToDeprecatedResponse Response Object
type UpdateMyActionTemplateToDeprecatedResponse struct {
	XRequestId *string `json:"x-request-id,omitempty"`

	ContentLength  *string `json:"Content-Length,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMyActionTemplateToDeprecatedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMyActionTemplateToDeprecatedResponse struct{}"
	}

	return strings.Join([]string{"UpdateMyActionTemplateToDeprecatedResponse", string(data)}, " ")
}
