package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLtsConfigsResponse Response Object
type ShowLtsConfigsResponse struct {

	// 实例LTS日志配置列表
	InstanceLtsConfigs *[]LtsConfigsV3 `json:"instance_lts_configs,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowLtsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLtsConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowLtsConfigsResponse", string(data)}, " ")
}
