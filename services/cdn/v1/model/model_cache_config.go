package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CacheConfig struct {
	// 是否忽略url中的参数。

	IgnoreUrlParameter *bool `json:"ignore_url_parameter,omitempty"`

	// 忽略的参数列表

	CacheUrlParamsConfig  *CacheUrlParamsConfig `json:"cache_url_params_config,omitempty"`

	// 缓存规则是否遵循源站

	FollowOrigin *bool `json:"follow_origin,omitempty"`

	Compress *CompressResponse `json:"compress,omitempty"`
	// 缓存规则，将覆盖之前的规则配置。规则为空重置为默认规则。

	Rules *[]Rules `json:"rules,omitempty"`
}

type CacheUrlParamsConfig struct {
	ParamType string `json:"param_type"` // reserve_params/del_params/"",表示保留参数，忽略参数，忽略所有
	ParamValue string `json:"param_value"`
}

func (o CacheConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheConfig struct{}"
	}

	return strings.Join([]string{"CacheConfig", string(data)}, " ")
}
