package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteMyActionTemplateResponse struct {
	XRequestId *string `json:"x-request-id,omitempty"`

	ContentLength  *string `json:"Content-Length,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMyActionTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMyActionTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteMyActionTemplateResponse", string(data)}, " ")
}
