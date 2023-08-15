package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndicatorsResponse Response Object
type ListIndicatorsResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// tatal count
	Total *int32 `json:"total,omitempty"`

	// list of informations of indicator
	Data *[]IndicatorDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIndicatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndicatorsResponse struct{}"
	}

	return strings.Join([]string{"ListIndicatorsResponse", string(data)}, " ")
}
