package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FeatureToggle struct {

	// 特性名称
	Name string `json:"name" xml:"name"`

	// 是否开启特性
	Enable bool `json:"enable" xml:"enable"`

	// 特性参数配置
	Config *string `json:"config,omitempty" xml:"config"`
}

func (o FeatureToggle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeatureToggle struct{}"
	}

	return strings.Join([]string{"FeatureToggle", string(data)}, " ")
}
