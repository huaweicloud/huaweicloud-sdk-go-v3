package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceResponse 资源信息
type ResourceResponse struct {

	// **参数解释：** 资源id **取值范围：** 不涉及
	ResourceId *string `json:"resourceId,omitempty"`

	// **参数解释：** 云服务产品对应的云服务类型 **取值范围：** 不涉及
	CloudServiceType *string `json:"cloudServiceType,omitempty"`

	// **参数解释：** 云服务产品的资源类型 **取值范围：**  - hws.resource.type.waf：云模式包周期WAF  - hws.resource.type.waf.domain：云模式包周期WAF域名扩展包  - hws.resource.type.waf.bandwidth：云模式包周期WAF带宽扩展包  - hws.resource.type.waf.rule：云模式包周期WAF规则扩展包  - hws.resource.type.waf.payperuserequest：Web应用防火墙按需请求   - hws.resource.type.waf.payperusedomain：Web应用防火墙按需域名  - hws.resource.type.waf.payperuserule：Web应用防火墙按需规则
	ResourceType *string `json:"resourceType,omitempty"`

	// **参数解释：** 云服务产品的资源规格 **取值范围：** 不涉及
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	// **参数解释：** 资源状态 **取值范围：**  - 0：解冻/正常   - 1：冻结   - 2：删除
	Status *int32 `json:"status,omitempty"`

	// **参数解释：** 资源到期时间 **取值范围：** 不涉及
	ExpireTime *string `json:"expireTime,omitempty"`

	// **参数解释：** 资源数量 **取值范围：** 不涉及
	ResourceSize *int32 `json:"resourceSize,omitempty"`

	// **参数解释：** 资源创建时间 **取值范围：** 不涉及
	Timestamp *int64 `json:"timestamp,omitempty"`

	// **参数解释：** 当资源项为附属资源(域名扩展包/带宽扩展包/规则扩展包/高阶功能扩展包),描述附属资源与主资源的关系 **取值范围：**  - 0：附属资源与主资源为挂载关系  - 1：附属资源与主资源为绑定关系  - 2：从属关系
	RelativeType *int32 `json:"relativeType,omitempty"`

	// **参数解释：** 华为云区域ID,资源所属的区域id **取值范围：** 不涉及
	Region *string `json:"region,omitempty"`
}

func (o ResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceResponse struct{}"
	}

	return strings.Join([]string{"ResourceResponse", string(data)}, " ")
}
