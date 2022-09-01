package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeLoadbalancerChargeModeResponse struct {

	// 转包周期下单成功的EIP ID列表
	EipIdList *[]string `json:"eip_id_list,omitempty" xml:"eip_id_list"`

	// 转包周期下单成功的LB ID列表
	LoadbalancerIdList *[]string `json:"loadbalancer_id_list,omitempty" xml:"loadbalancer_id_list"`

	// 转包周期订单号
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 请求的UUIID
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeLoadbalancerChargeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerChargeModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerChargeModeResponse", string(data)}, " ")
}
