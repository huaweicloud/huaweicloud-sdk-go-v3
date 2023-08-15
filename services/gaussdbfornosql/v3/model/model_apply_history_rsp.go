package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyHistoryRsp struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 实例名称
	InstanceName string `json:"instance_name"`

	// 生效时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。  [其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。](tag:hc)  [其中，T指某个时间的开始；Z指时区偏移量。](tag:hk)
	AppliedAt *sdktime.SdkTime `json:"applied_at"`

	// - SUCCESS：应用成功。 - FAILED:应用失败。
	ApplyResult string `json:"apply_result"`

	// 失败原因
	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o ApplyHistoryRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyHistoryRsp struct{}"
	}

	return strings.Join([]string{"ApplyHistoryRsp", string(data)}, " ")
}
