package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeInstanceResponse Response Object
type ResizeInstanceResponse struct {

	// **参数说明**：订单号，修改包年包月实例时返回该参数，修改按需实例时返回为空。查看订单详情请参考[[查询订单详情](https://support.huaweicloud.com/api-bpconsole/zh-cn_topic_0075746564.html)]\"
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeInstanceResponse", string(data)}, " ")
}
