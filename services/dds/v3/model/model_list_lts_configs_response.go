package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsConfigsResponse Response Object
type ListLtsConfigsResponse struct {

	// 每个实例的LTS日志配置信息和实例简要信息。
	InstanceLtsConfigs *[]ListLtsLogPolicyRespondBodyInstanceLtsConfigs `json:"instance_lts_configs,omitempty"`

	// 全部可查询的云服务日志配置结果个数，等于所有DDS实例的个数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLtsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsConfigsResponse", string(data)}, " ")
}
