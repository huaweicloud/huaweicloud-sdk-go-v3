package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListThumbnailDetailsRequest Request Object
type ListThumbnailDetailsRequest struct {

	// 截图对应的任务id
	TaskId string `json:"task_id"`

	// 查询偏移量。取值范围[0,20480]，默认值：0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询一页返回数。取值范围[1,100]，默认值：20。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListThumbnailDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListThumbnailDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListThumbnailDetailsRequest", string(data)}, " ")
}
