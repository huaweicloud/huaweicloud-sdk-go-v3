package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRapidGrowthThresholdNewResponse Response Object
type SetRapidGrowthThresholdNewResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SetRapidGrowthThresholdNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRapidGrowthThresholdNewResponse struct{}"
	}

	return strings.Join([]string{"SetRapidGrowthThresholdNewResponse", string(data)}, " ")
}
