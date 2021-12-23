package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSpecifiedApiVersionRequest struct {
	// API版本号。例如: v1。

	ApiVersion string `json:"api_version"`
}

func (o ShowSpecifiedApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpecifiedApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowSpecifiedApiVersionRequest", string(data)}, " ")
}
