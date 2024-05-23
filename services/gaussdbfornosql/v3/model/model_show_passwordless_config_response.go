package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPasswordlessConfigResponse Response Object
type ShowPasswordlessConfigResponse struct {

	// 免密配置,IP与网段的列表,仅支持ipv4的IP或网段。
	ConfigIps *[]string `json:"config_ips,omitempty"`

	// 免密配置的总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPasswordlessConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPasswordlessConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowPasswordlessConfigResponse", string(data)}, " ")
}
