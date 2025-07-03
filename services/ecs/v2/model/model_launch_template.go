package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LaunchTemplate struct {

	// 模板id
	Id string `json:"id"`

	// 模板名称
	Name string `json:"name"`

	// 模板描述
	Description string `json:"description"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`

	// 模板默认版本号
	DefaultVersion int32 `json:"default_version"`

	// 模板最新版本号
	LatestVersion int32 `json:"latest_version"`
}

func (o LaunchTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LaunchTemplate struct{}"
	}

	return strings.Join([]string{"LaunchTemplate", string(data)}, " ")
}
