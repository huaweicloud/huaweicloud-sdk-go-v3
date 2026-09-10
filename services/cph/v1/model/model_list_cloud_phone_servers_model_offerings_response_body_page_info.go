package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneServersModelOfferingsResponseBodyPageInfo 页标记。
type ListCloudPhoneServersModelOfferingsResponseBodyPageInfo struct {

	// 返回下一页查询地址。
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o ListCloudPhoneServersModelOfferingsResponseBodyPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServersModelOfferingsResponseBodyPageInfo struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServersModelOfferingsResponseBodyPageInfo", string(data)}, " ")
}
