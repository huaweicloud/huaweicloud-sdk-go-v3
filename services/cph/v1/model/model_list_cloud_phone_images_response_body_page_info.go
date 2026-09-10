package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneImagesResponseBodyPageInfo 页标记。
type ListCloudPhoneImagesResponseBodyPageInfo struct {

	// 返回下一页查询marker。
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o ListCloudPhoneImagesResponseBodyPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneImagesResponseBodyPageInfo struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneImagesResponseBodyPageInfo", string(data)}, " ")
}
