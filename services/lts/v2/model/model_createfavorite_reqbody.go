package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatefavoriteReqbody struct {

	// 企业项目id
	EpsId *string `json:"eps_id,omitempty"`

	// 收藏资源id
	FavoriteResourceId string `json:"favorite_resource_id"`

	// 收藏资源类型
	FavoriteResourceType string `json:"favorite_resource_type"`

	// 日志组id
	LogGroupId string `json:"log_group_id"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 日志流id
	LogStreamId string `json:"log_stream_id"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 是否支持全局化，必填true，否则创建不了收藏
	IsGlobal bool `json:"is_global"`
}

func (o CreatefavoriteReqbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatefavoriteReqbody struct{}"
	}

	return strings.Join([]string{"CreatefavoriteReqbody", string(data)}, " ")
}
