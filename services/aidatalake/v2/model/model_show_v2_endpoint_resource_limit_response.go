package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2EndpointResourceLimitResponse Response Object
type ShowV2EndpointResourceLimitResponse struct {
	Body *[]ResourceLimit `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowV2EndpointResourceLimitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2EndpointResourceLimitResponse struct{}"
	}

	return strings.Join([]string{"ShowV2EndpointResourceLimitResponse", string(data)}, " ")
}
