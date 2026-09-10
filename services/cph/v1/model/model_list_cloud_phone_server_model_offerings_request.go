package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneServerModelOfferingsRequest Request Object
type ListCloudPhoneServerModelOfferingsRequest struct {

	// 分页标记。从marker指定的下一条数据开始查询。
	Marker *string `json:"marker,omitempty"`

	// 最小值1，最大值1000，默认为100。返回的结果中记录数不超过limit值。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCloudPhoneServerModelOfferingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServerModelOfferingsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServerModelOfferingsRequest", string(data)}, " ")
}
