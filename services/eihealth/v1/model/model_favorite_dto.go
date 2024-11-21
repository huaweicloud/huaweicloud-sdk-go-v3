package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FavoriteDto struct {

	// 收藏ID。
	Id *string `json:"id,omitempty"`

	// 收藏类型。
	Type *string `json:"type,omitempty"`

	// 收藏者的用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 收藏者的用户名称。
	UserName *string `json:"user_name,omitempty"`

	// 收藏时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 展示信息。
	DisplayInfo *string `json:"display_info,omitempty"`

	// 定位信息。
	LocationInfo *string `json:"location_info,omitempty"`
}

func (o FavoriteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FavoriteDto struct{}"
	}

	return strings.Join([]string{"FavoriteDto", string(data)}, " ")
}
