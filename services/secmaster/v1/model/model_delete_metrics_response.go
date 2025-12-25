package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMetricsResponse Response Object
type DeleteMetricsResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMetricsResponse struct{}"
	}

	return strings.Join([]string{"DeleteMetricsResponse", string(data)}, " ")
}
