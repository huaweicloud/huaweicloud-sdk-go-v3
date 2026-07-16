package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateIntranetConnectionRequestBody struct {

	// **参数解释：** 内网接入申请操作。 **约束限制：** 不涉及。 **取值范围：** - APPROVE：通过申请。只有当内网接入申请处于“审批中”状态时，才可进行此操作。 - REJECT： 拒绝申请。只有当内网接入申请处于“审批中”状态时，才可进行此操作。 - CANCEL： 取消授权，只有当内网接入申请处于“通过”（CONNECTED）状态时，才可进行取消授权操作。 - RETRY：  重试申请，只有当内网接入申请处于“异常”状态并且异常原因为“连接失败，请重试”时，才可进行重试操作。 **默认取值：** 不涉及。
	Action string `json:"action"`

	// **参数解释：** 拒绝时可以填入拒绝的原因。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Reason *string `json:"reason,omitempty"`
}

func (o UpdateIntranetConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIntranetConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIntranetConnectionRequestBody", string(data)}, " ")
}
