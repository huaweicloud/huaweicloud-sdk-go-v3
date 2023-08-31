package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetPrimaryVideoThumbnailRequestBody 设置视频模板封面图请求体。
type SetPrimaryVideoThumbnailRequestBody struct {

	// AIM资源ID。
	AimResourceId string `json:"aim_resource_id"`

	// 视频封面图ID。
	ThumbnailId string `json:"thumbnail_id"`
}

func (o SetPrimaryVideoThumbnailRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPrimaryVideoThumbnailRequestBody struct{}"
	}

	return strings.Join([]string{"SetPrimaryVideoThumbnailRequestBody", string(data)}, " ")
}
