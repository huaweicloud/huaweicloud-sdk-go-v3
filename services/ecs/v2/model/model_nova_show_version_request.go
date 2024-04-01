package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaShowVersionRequest Request Object
type NovaShowVersionRequest struct {

	// API版本号。例如: v2
	ApiVersion string `json:"api_version"`
}

func (o NovaShowVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowVersionRequest struct{}"
	}

	return strings.Join([]string{"NovaShowVersionRequest", string(data)}, " ")
}
