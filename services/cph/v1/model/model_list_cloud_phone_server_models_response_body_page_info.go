package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneServerModelsResponseBodyPageInfo 页标记。
type ListCloudPhoneServerModelsResponseBodyPageInfo struct {

	// 返回下一页查询marker地址。
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o ListCloudPhoneServerModelsResponseBodyPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServerModelsResponseBodyPageInfo struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServerModelsResponseBodyPageInfo", string(data)}, " ")
}
