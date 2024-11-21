package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFavoriteReq 添加收藏请求体。
type CreateFavoriteReq struct {

	// 收藏类型。
	Type string `json:"type"`

	// 收藏的资源ID。
	ResourceId string `json:"resource_id"`

	// 收藏的资源名称，正则匹配中文，英文字母和数字及下划线。
	ResourceName string `json:"resource_name"`

	// 收藏的资源类型，正则匹配英文字母和数字及下划线。
	ResourceType string `json:"resource_type"`

	// 展示信息。
	DisplayInfo string `json:"display_info"`

	// 定位信息。
	LocationInfo string `json:"location_info"`
}

func (o CreateFavoriteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFavoriteReq struct{}"
	}

	return strings.Join([]string{"CreateFavoriteReq", string(data)}, " ")
}
