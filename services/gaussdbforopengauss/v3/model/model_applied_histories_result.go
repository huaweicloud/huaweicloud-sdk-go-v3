package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppliedHistoriesResult struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 应用状态 (SUCCESS | FAILED)。
	ApplyResult string `json:"apply_result"`

	// 应用时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	AppliedAt string `json:"applied_at"`

	// 失败原因错误码，如DBS.280005。
	ErrorCode string `json:"error_code"`
}

func (o AppliedHistoriesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppliedHistoriesResult struct{}"
	}

	return strings.Join([]string{"AppliedHistoriesResult", string(data)}, " ")
}
