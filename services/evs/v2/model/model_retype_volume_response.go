package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetypeVolumeResponse Response Object
type RetypeVolumeResponse struct {

	// 任务ID，云硬盘为按需计费时返回该参数。 > 说明： >  > 如果需要查询job的状态，请参考：\"[查询job的状态](https://support.huaweicloud.com/api-evs/evs_04_0054.html)\"。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，云硬盘为包周期计费时返回该参数。 > 说明： >  > - 如果您需要支付订单，请参考：\"[支付包周期产品订单](https://support.huaweicloud.com/api-oce/zh-cn_topic_0075746561.html)\"。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RetypeVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetypeVolumeResponse struct{}"
	}

	return strings.Join([]string{"RetypeVolumeResponse", string(data)}, " ")
}
