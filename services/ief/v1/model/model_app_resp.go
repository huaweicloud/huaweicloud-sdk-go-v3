package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用模板配置
type AppResp struct {
	// 应用模板ID

	Id string `json:"id"`
	// 应用模板名称，只允许英文小写字母、数字、中划线，最大长度32, 英文小写字母或数字开头和结尾 Name为必填字段，且本租户中唯一

	Name string `json:"name"`
	// 应用模板别名，中文英文字母、数字、中划线、下划线，最大64字符

	Alias string `json:"alias"`
	// 应用模板描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\

	Description string `json:"description"`
	// 应用图标存储url地址，最大长度2083

	IconUrl string `json:"icon_url"`
	// 创建时间

	CreatedAt string `json:"created_at"`
	// 更新时间

	UpdatedAt string `json:"updated_at"`
	// 项目ID

	ProjectId string `json:"project_id"`
	// 模板类型

	Visibility string `json:"visibility"`
	// app详情

	AppVersions []AppVersionDetail `json:"app_versions"`
}

func (o AppResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppResp struct{}"
	}

	return strings.Join([]string{"AppResp", string(data)}, " ")
}
