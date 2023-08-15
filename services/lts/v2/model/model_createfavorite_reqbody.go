package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

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
}

func (o CreatefavoriteReqbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatefavoriteReqbody struct{}"
	}

	return strings.Join([]string{"CreatefavoriteReqbody", string(data)}, " ")
}
