package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelServiceConfig 服务配置
type ModelServiceConfig struct {
	CustomSpec *ModelServiceConfigCustomSpec `json:"custom_spec,omitempty"`

	// 部署环境变量，Map格式
	UserEnv map[string]string `json:"user_env,omitempty"`

	// 部署实例个数，取值范围[1-10]，默认值1
	InstanceCount int32 `json:"instance_count"`

	// 请求大小限制，取值范围[1-64000]，默认值6400
	RequestSizeLimit *int32 `json:"request_size_limit,omitempty"`

	// 请求QPS限制，取值范围[1-10000]，默认值100
	RequestLimitPerSecond *int32 `json:"request_limit_per_second,omitempty"`

	// 请求超时时间，取值范围[1-120000]，默认值120000，单位毫秒
	RequestTimeout *int32 `json:"request_timeout,omitempty"`
}

func (o ModelServiceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelServiceConfig struct{}"
	}

	return strings.Join([]string{"ModelServiceConfig", string(data)}, " ")
}
