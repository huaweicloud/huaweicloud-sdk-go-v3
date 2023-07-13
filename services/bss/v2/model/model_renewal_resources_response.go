package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenewalResourcesResponse Response Object
type RenewalResourcesResponse struct {

	// 续订资源生成的订单ID的列表。
	OrderIds *[]string `json:"order_ids,omitempty"`

	// |参数名称：失败的资源信息列表| |参数的约束及描述：套餐包使用量信息|
	FailResourceInfos *[]FailResourceInfo `json:"fail_resource_infos,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o RenewalResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenewalResourcesResponse struct{}"
	}

	return strings.Join([]string{"RenewalResourcesResponse", string(data)}, " ")
}
