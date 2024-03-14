package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionObject 描述VPCEP服务API版本信息列表
type VersionObject struct {

	// 版本状态。   - CURRENT：表示该版本为主推版本。   - SUPPORT：表示为老版本，但是现在还在继续支持。   -DEPRECATED：表示为废弃版本，存在后续删除的可能。
	Status *string `json:"status,omitempty"`

	// 版本ID。 - v1：当前主推版本
	Id *string `json:"id,omitempty"`

	// 版本发布时间。采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ
	Updated *sdktime.SdkTime `json:"updated,omitempty"`

	// 支持的版本号。
	Version *string `json:"version,omitempty"`

	// 支持的微版本号。若该版本API不支持微版本，则为空。
	MinVersion *string `json:"min_version,omitempty"`

	// API的URL地址
	Links *[]Link `json:"links,omitempty"`
}

func (o VersionObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionObject struct{}"
	}

	return strings.Join([]string{"VersionObject", string(data)}, " ")
}
