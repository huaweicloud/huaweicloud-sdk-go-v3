package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMetricsConfigResponse Response Object
type UpdateMetricsConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMetricsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMetricsConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateMetricsConfigResponse", string(data)}, " ")
}
