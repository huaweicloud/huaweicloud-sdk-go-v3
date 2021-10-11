package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

type FeatureInfo struct {
	// 特性编号

	Id *string `json:"id,omitempty"`
	// 特性名称

	Name *string `json:"name,omitempty"`
	// 是否开启特性

	Enable *bool `json:"enable,omitempty"`
	// 特性参数配置

	Config *string `json:"config,omitempty"`
	// 实例编号

	InstanceId *string `json:"instance_id,omitempty"`
	// 实例特性更新时间

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o FeatureInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FeatureInfo struct{}"
	}

	return strings.Join([]string{"FeatureInfo", string(data)}, " ")
}
