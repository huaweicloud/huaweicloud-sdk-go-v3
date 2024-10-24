package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceUsageResponse Response Object
type ShowResourceUsageResponse struct {

	// 资源用量列表
	Resources *[]ResourceUsageInfo `json:"resources,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResourceUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceUsageResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceUsageResponse", string(data)}, " ")
}
