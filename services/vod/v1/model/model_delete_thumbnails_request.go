package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteThumbnailsRequest Request Object
type DeleteThumbnailsRequest struct {

	// 媒资ID
	AssetId *string `json:"asset_id,omitempty"`

	// 删除指定截图任务的截图结果，一次最多10个
	TaskId *[]string `json:"task_id,omitempty"`
}

func (o DeleteThumbnailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteThumbnailsRequest struct{}"
	}

	return strings.Join([]string{"DeleteThumbnailsRequest", string(data)}, " ")
}
