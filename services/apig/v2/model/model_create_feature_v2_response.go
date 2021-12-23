package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateFeatureV2Response struct {
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

	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateFeatureV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFeatureV2Response struct{}"
	}

	return strings.Join([]string{"CreateFeatureV2Response", string(data)}, " ")
}
