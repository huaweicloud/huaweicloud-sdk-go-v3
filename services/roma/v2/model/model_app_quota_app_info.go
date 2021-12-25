package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppQuotaAppInfo struct {
	// 客户端应用编号

	AppId *string `json:"app_id,omitempty"`
	// 客户端应用名称

	Name *string `json:"name,omitempty"`
	// 客户端应用状态： - 1：启用 - 2：禁用

	Status *int32 `json:"status,omitempty"`
	// 客户端应用的Key

	AppKey *string `json:"app_key,omitempty"`
	// 客户端应用描述

	Remark *string `json:"remark,omitempty"`
	// 创建时间

	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`
	// 更新时间

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
	// 客户端配额编号

	AppQuotaId *string `json:"app_quota_id,omitempty"`
	// 配额名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3-255字符

	AppQuotaName *string `json:"app_quota_name,omitempty"`
	// 绑定时间

	BoundTime *sdktime.SdkTime `json:"bound_time,omitempty"`
}

func (o AppQuotaAppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppQuotaAppInfo struct{}"
	}

	return strings.Join([]string{"AppQuotaAppInfo", string(data)}, " ")
}
