package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceChargeModeResponse Response Object
type ChangeInstanceChargeModeResponse struct {

	// **参数说明**：订单号。[查看订单详情请参考[查询订单详情](https://support.huaweicloud.com/api-bpconsole/zh-cn_topic_0075746564.html)。](tag:hws)\"
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeInstanceChargeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceChargeModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeInstanceChargeModeResponse", string(data)}, " ")
}
