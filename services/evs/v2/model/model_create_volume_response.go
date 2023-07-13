package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVolumeResponse Response Object
type CreateVolumeResponse struct {

	// 任务ID，云硬盘为按需计费时返回该参数。 > 说明： >  > 如果需要查询job的状态，请参考：\"[查询job的状态](https://support.huaweicloud.com/api-evs/evs_04_0054.html)\"。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，云硬盘为包周期计费时返回该参数。 > 说明： > 直接在包周期云服务器上新增云硬盘，系统会自动将云硬盘挂载到包周期云服务器上。该情形下也会返回该参数。  > - 如果您需要支付订单，请参考：\"[支付包周期产品订单](https://support.huaweicloud.com/api-oce/zh-cn_topic_0075746561.html)\"。
	OrderId *string `json:"order_id,omitempty"`

	// 待创建的云硬盘ID列表。 > 说明： > 通过云硬盘ID查询云硬盘详情 ，若返回404 可能云硬盘正在创建中或者已经创建失败。 > 通过JobId查询云硬盘创建任务是否完成[查询job的状态](https://support.huaweicloud.com/api-evs/evs_04_0054.html)。
	VolumeIds      *[]string `json:"volume_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeResponse struct{}"
	}

	return strings.Join([]string{"CreateVolumeResponse", string(data)}, " ")
}
