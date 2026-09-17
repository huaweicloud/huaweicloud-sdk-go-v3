package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportedMetricsResponse Response Object
type ListSupportedMetricsResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSupportedMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListSupportedMetricsResponse", string(data)}, " ")
}
