package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFsDirUsageResponse Response Object
type ShowFsDirUsageResponse struct {
	DirUsage *FsDirUasge `json:"dir_usage,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFsDirUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFsDirUsageResponse struct{}"
	}

	return strings.Join([]string{"ShowFsDirUsageResponse", string(data)}, " ")
}
