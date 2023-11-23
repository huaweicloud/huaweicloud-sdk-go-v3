package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsConfigsResponse Response Object
type ListLtsConfigsResponse struct {

	// 实例总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 实例-LTS日志配置信息列表。
	InstanceLtsConfigs *[]InstanceLogConfig `json:"instance_lts_configs,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListLtsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsConfigsResponse", string(data)}, " ")
}
