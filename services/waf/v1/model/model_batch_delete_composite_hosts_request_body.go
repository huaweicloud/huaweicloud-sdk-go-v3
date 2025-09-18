package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteCompositeHostsRequestBody struct {

	// **参数解释：** 域名id数组，包含待批量删除的租户域名唯一标识，从查询租户域名列表接口获取。 **约束限制：** 不涉及 **取值范围：** 数组内元素为字符串类型，每个元素对应一个域名ID **默认取值：** 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释：** 保留域名的防护策略，标识删除域名后是否保留其关联的防护策略（true表示保留，false表示不保留）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** false
	KeepPolicyEnable *bool `json:"keep_policy_enable,omitempty"`
}

func (o BatchDeleteCompositeHostsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCompositeHostsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteCompositeHostsRequestBody", string(data)}, " ")
}
