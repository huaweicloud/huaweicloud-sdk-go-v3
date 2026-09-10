package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCloudPhoneServersModelOfferingsResponseBodyModels struct {

	// 可用区
	AvailableZone *string `json:"available_zone,omitempty"`

	// 云手机服务器规格名
	ModelName *string `json:"model_name,omitempty"`

	// 云手机服务器规格的售卖状态。  - sellout：售罄  - available：可用
	SellStatus *string `json:"sell_status,omitempty"`
}

func (o ListCloudPhoneServersModelOfferingsResponseBodyModels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServersModelOfferingsResponseBodyModels struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServersModelOfferingsResponseBodyModels", string(data)}, " ")
}
