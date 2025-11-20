package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Instance struct {

	// - 参数解释：ESW实例的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Id string `json:"id"`

	// - 参数解释：ESW实例名称。 - 约束限制：   - 长度范围为1~64个字符。   - 名称由中文、英文字母、数字、下划线（_）、中划线（-）、点（.）组成。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name string `json:"name"`

	// - 参数解释：ESW实例所属项目ID。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ProjectId string `json:"project_id"`

	// - 参数解释：ESW实例所属region的ID。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Region string `json:"region"`

	// - 参数解释：ESW实例规格。 - 约束限制：不涉及。 - 取值范围：参考flavor list接口响应。 - 默认取值：不涉及。
	FlavorRef string `json:"flavor_ref"`

	// - 参数解释：ESW实例的高可用模式。 - 约束限制：当前只支持设置ha。 - 取值范围：ha。 - 默认取值：不涉及。
	HaMode string `json:"ha_mode"`

	// - 参数解释：ESW实例的状态。 - 约束限制：不涉及。 - 取值范围：   - active：运行中   - failed：创建失败   - abnormal：异常   - build：创建中   - frozen：已冻结 - 默认取值：不涉及。
	Status string `json:"status"`

	// - 参数解释：ESW实例创建时间。 - 约束限制：UTC时间，格式为yyyy-MM-ddTHH:mm:ss。 - 取值范围：不涉及。 - 默认取值：不涉及。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// - 参数解释：ESW实例更新时间。 - 约束限制：UTC时间，格式为yyyy-MM-ddTHH:mm:ss。 - 取值范围：不涉及。 - 默认取值：不涉及。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// - 参数解释：ESW实例描述信息。 - 约束限制：   - 长度范围为0~255个字符。   - 不能包含“<”和“>”。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Description string `json:"description"`

	TunnelInfo *TunnelInfoResult `json:"tunnel_info"`

	ChargeInfos *PostPaidChargeInfos `json:"charge_infos"`

	AvailabilityZones *AvailabilityZones `json:"availability_zones"`
}

func (o Instance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instance struct{}"
	}

	return strings.Join([]string{"Instance", string(data)}, " ")
}
