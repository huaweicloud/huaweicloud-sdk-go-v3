package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppIdInfoDto 添加企业应用的返回结果
type AppIdInfoDto struct {

	// 企业应用名称
	AppName *string `json:"app_name,omitempty"`

	// 企业应用
	AppId *string `json:"app_id,omitempty"`

	// 企业应用appkey
	AppKey *string `json:"app_key,omitempty"`

	// 企业应用描述
	Description *string `json:"description,omitempty"`

	// 企业应用创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最近修改时间
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 企业应用状态  * 0：正常  * 1：停用
	Status *int32 `json:"status,omitempty"`
}

func (o AppIdInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppIdInfoDto struct{}"
	}

	return strings.Join([]string{"AppIdInfoDto", string(data)}, " ")
}
