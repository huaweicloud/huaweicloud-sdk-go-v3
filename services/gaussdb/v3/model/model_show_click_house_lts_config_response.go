package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClickHouseLtsConfigResponse Response Object
type ShowClickHouseLtsConfigResponse struct {

	// 实例LTS配置信息。
	InstanceLtsConfigs *[]ChInstanceLtsConfigs `json:"instance_lts_configs,omitempty"`

	// 实例总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowClickHouseLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClickHouseLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowClickHouseLtsConfigResponse", string(data)}, " ")
}
