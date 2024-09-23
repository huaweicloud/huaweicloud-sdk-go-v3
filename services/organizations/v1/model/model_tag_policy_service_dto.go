package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagPolicyServiceDto 被添加到标签策略强制执行的资源类型。
type TagPolicyServiceDto struct {

	// 服务名称。
	ServiceName string `json:"service_name"`

	// 资源类型。
	ResourceTypes []string `json:"resource_types"`

	// resource_type是否支持全量选择，即*
	SupportAll bool `json:"support_all"`
}

func (o TagPolicyServiceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagPolicyServiceDto struct{}"
	}

	return strings.Join([]string{"TagPolicyServiceDto", string(data)}, " ")
}
