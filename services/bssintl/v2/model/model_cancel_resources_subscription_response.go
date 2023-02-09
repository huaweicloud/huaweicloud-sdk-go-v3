package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CancelResourcesSubscriptionResponse struct {

	// 客户退订订单ID的列表信息。
	OrderIds *[]string `json:"order_ids,omitempty"`

	// |参数名称：失败的资源信息列表| |参数的约束及描述：套餐包使用量信息|
	FailResourceInfos *[]FailResourceInfo `json:"fail_resource_infos,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o CancelResourcesSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelResourcesSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"CancelResourcesSubscriptionResponse", string(data)}, " ")
}
