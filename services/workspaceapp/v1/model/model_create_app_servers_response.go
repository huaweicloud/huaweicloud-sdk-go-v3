package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppServersResponse Response Object
type CreateAppServersResponse struct {

	// 对于创建云应用服务器命令下发后会返回job_id，通过job_id可以查询任务的执行状态,按需购买返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 订单号，创建包年包月的弹性云服务器时返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppServersResponse struct{}"
	}

	return strings.Join([]string{"CreateAppServersResponse", string(data)}, " ")
}
