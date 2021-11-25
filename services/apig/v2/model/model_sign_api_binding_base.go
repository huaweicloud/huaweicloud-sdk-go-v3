package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SignApiBindingBase struct {
	// API的发布编号

	PublishId *string `json:"publish_id,omitempty"`
	// API编号

	ApiId *string `json:"api_id,omitempty"`
	// API所属分组的名称

	GroupName *string `json:"group_name,omitempty"`
	// 绑定时间

	BindingTime *sdktime.SdkTime `json:"binding_time,omitempty"`
	// API所属环境的编号

	EnvId *string `json:"env_id,omitempty"`
	// API所属环境的名称

	EnvName *string `json:"env_name,omitempty"`
	// API类型

	ApiType *int32 `json:"api_type,omitempty"`
	// API名称

	ApiName *string `json:"api_name,omitempty"`
	// 绑定关系的ID

	Id *string `json:"id,omitempty"`
	// API描述

	ApiRemark *string `json:"api_remark,omitempty"`
	// 签名密钥的编号

	SignId *string `json:"sign_id,omitempty"`
	// 签名密钥的名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。

	SignName *string `json:"sign_name,omitempty"`
}

func (o SignApiBindingBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignApiBindingBase struct{}"
	}

	return strings.Join([]string{"SignApiBindingBase", string(data)}, " ")
}
