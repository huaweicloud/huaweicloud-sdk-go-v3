package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneServerModelsRequest Request Object
type ListCloudPhoneServerModelsRequest struct {

	// 产品类型。 - 0：云手机 - 1：云手游
	ProductType *int32 `json:"product_type,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的服务器规格个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCloudPhoneServerModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServerModelsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServerModelsRequest", string(data)}, " ")
}
