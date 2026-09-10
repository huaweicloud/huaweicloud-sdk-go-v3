package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneServerModelOfferingsResponse Response Object
type ListCloudPhoneServerModelOfferingsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 规格总数。
	Count *int32 `json:"count,omitempty"`

	// 云手机服务器规格信息
	Models *[]ListCloudPhoneServersModelOfferingsResponseBodyModels `json:"models,omitempty"`

	PageInfo       *ListCloudPhoneServersModelOfferingsResponseBodyPageInfo `json:"page_info,omitempty"`
	HttpStatusCode int                                                      `json:"-"`
}

func (o ListCloudPhoneServerModelOfferingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServerModelOfferingsResponse struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServerModelOfferingsResponse", string(data)}, " ")
}
