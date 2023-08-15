package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorResponse Response Object
type CreateIndicatorResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *IndicatorDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIndicatorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorResponse struct{}"
	}

	return strings.Join([]string{"CreateIndicatorResponse", string(data)}, " ")
}
