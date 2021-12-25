package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppConfigInfoV2 struct {
	// 应用配置编号

	Id *string `json:"id,omitempty"`
	// 应用编号

	AppId *string `json:"app_id,omitempty"`
	// 应用配置类型

	ConfigType *string `json:"config_type,omitempty"`
	// 应用配置名称

	ConfigName *string `json:"config_name,omitempty"`
	// 应用配置值

	ConfigValue *string `json:"config_value,omitempty"`
	// 应用配置更新时间

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
	// 应用配置描述

	Description *string `json:"description,omitempty"`
}

func (o AppConfigInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppConfigInfoV2 struct{}"
	}

	return strings.Join([]string{"AppConfigInfoV2", string(data)}, " ")
}
