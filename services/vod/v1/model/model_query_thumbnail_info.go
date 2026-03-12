package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryThumbnailInfo 截图信息。
type QueryThumbnailInfo struct {

	// 截图任务
	TaskId *string `json:"task_id,omitempty"`

	ThumbnailPara *Thumbnail `json:"thumbnail_para,omitempty"`

	// 视频截图个数。
	ThumbnailCount *int32 `json:"thumbnail_count,omitempty"`
}

func (o QueryThumbnailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryThumbnailInfo struct{}"
	}

	return strings.Join([]string{"QueryThumbnailInfo", string(data)}, " ")
}
