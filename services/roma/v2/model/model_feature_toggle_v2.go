package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FeatureToggleV2 struct {

	// 特性名称
	Name string `json:"name" xml:"name"`

	// 是否开启特性
	Enable bool `json:"enable" xml:"enable"`

	// 特性参数配置
	Config *string `json:"config,omitempty" xml:"config"`
}

func (o FeatureToggleV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeatureToggleV2 struct{}"
	}

	return strings.Join([]string{"FeatureToggleV2", string(data)}, " ")
}
