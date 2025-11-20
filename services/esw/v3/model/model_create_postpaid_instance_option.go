package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePostpaidInstanceOption struct {
	AvailabilityZones *AvailabilityZones `json:"availability_zones"`

	ChargeInfos *PostPaidChargeInfos `json:"charge_infos"`

	// - 参数解释：ESW实例规格。 - 约束限制：不涉及。 - 取值范围：参考flavor list接口响应。 - 默认取值：不涉及。
	FlavorRef string `json:"flavor_ref"`

	// - 参数解释：ESW实例的高可用模式。 - 约束限制：当前只支持设置ha。 - 取值范围：ha。 - 默认取值：不涉及。
	HaMode string `json:"ha_mode"`

	// - 参数解释：ESW实例名称。 - 约束限制：   - 长度范围为1~64个字符。   - 名称由中文、英文字母、数字、下划线（_）、中划线（-）、点（.）组成。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name string `json:"name"`

	TunnelInfo *TunnelInfoOption `json:"tunnel_info"`

	// - 参数解释：ESW实例描述信息。 - 约束限制：   - 长度范围为0~255个字符。   - 不能包含“<”和“>”。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Description *string `json:"description,omitempty"`
}

func (o CreatePostpaidInstanceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostpaidInstanceOption struct{}"
	}

	return strings.Join([]string{"CreatePostpaidInstanceOption", string(data)}, " ")
}
