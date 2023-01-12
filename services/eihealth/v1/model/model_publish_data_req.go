package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 发布数据资产请求体
type PublishDataReq struct {

	// 资产名
	Name string `json:"name"`

	// 版本号
	Version string `json:"version"`

	// 展示名
	Title *string `json:"title,omitempty"`

	// 短描述
	Summary *string `json:"summary,omitempty"`

	// 详细描述
	Description *string `json:"description,omitempty"`

	// 封面图片base64编码
	Picture *string `json:"picture,omitempty"`

	// 标签列表
	Labels *[]string `json:"labels,omitempty"`

	// 路径列表
	Paths []string `json:"paths"`
}

func (o PublishDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishDataReq struct{}"
	}

	return strings.Join([]string{"PublishDataReq", string(data)}, " ")
}
