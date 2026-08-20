package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsDetailsResponse Response Object
type ListEndpointsDetailsResponse struct {
	Error *ErrorInfo `json:"error,omitempty"`

	Result *EndpointList `json:"result,omitempty"`

	// 状态值
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEndpointsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointsDetailsResponse", string(data)}, " ")
}
