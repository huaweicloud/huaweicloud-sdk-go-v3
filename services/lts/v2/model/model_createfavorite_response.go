package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatefavoriteResponse struct {

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 企业项目id
	EpsId *string `json:"eps_id,omitempty"`

	// 收藏资源id
	FavoriteResourceId *string `json:"favorite_resource_id,omitempty"`

	// 收藏资源类型
	FavoriteResourceType *string `json:"favorite_resource_type,omitempty"`

	// 日志组id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 日志流id
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 项目id
	ProjectId      *string `json:"project_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatefavoriteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatefavoriteResponse struct{}"
	}

	return strings.Join([]string{"CreatefavoriteResponse", string(data)}, " ")
}
