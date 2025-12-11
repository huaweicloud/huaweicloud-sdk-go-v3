package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyHistoryRsp struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 实例名称。 **取值范围：** 不涉及。
	InstanceName string `json:"instance_name"`

	// **参数解释：** 生效时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 其中，T指某个时间的开始；Z指时区偏移量。 **取值范围：** 不涉及。
	AppliedAt *sdktime.SdkTime `json:"applied_at"`

	// **参数解释：** 应用结果。 **取值范围：** - SUCCESS：应用成功。 - FAILED:应用失败。 - APPLYING: 应用中。
	ApplyResult string `json:"apply_result"`

	// **参数解释：** 失败原因。 **取值范围：** 不涉及。
	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o ApplyHistoryRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyHistoryRsp struct{}"
	}

	return strings.Join([]string{"ApplyHistoryRsp", string(data)}, " ")
}
