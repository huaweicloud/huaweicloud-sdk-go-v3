package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListThumbnailInfoRequest Request Object
type ListThumbnailInfoRequest struct {

	// 截图对应媒资id的截图，只支持单个媒资。
	AssetId *string `json:"asset_id,omitempty"`

	// 截图对应的任务id，只支持单个任务查询。
	TaskId *string `json:"task_id,omitempty"`

	// 查询偏移量。取值范围[0,20000]，默认值：0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询一页返回数。取值范围[1,100]，默认值：10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListThumbnailInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListThumbnailInfoRequest struct{}"
	}

	return strings.Join([]string{"ListThumbnailInfoRequest", string(data)}, " ")
}
