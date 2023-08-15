package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneServerModelsRequest Request Object
type ListCloudPhoneServerModelsRequest struct {

	// 产品类型 - 0：云手机 - 1：云手游
	ProductType *int32 `json:"product_type,omitempty"`
}

func (o ListCloudPhoneServerModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServerModelsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServerModelsRequest", string(data)}, " ")
}
