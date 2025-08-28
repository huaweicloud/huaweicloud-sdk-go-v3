package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpgradeV3RequestBody **参数解释**：升级负载均衡器请求体。  **约束限制**：不涉及
type UpgradeV3RequestBody struct {

	// **参数解释**：升级过程的操作。  **约束限制**：不涉及  **取值范围**： - start：开始升级。只有当负载均衡器的provisioning_status为ACTIVE时，才能开始升级。 - complete：确认升级完成。只有当实例的provisioning_status为UPGRADED时，才能确认升级完成。确认后，实例无法再执行回退。 - rollback：回退升级，只有当实例的provisioning_status为UPGRADED、UPGRADE_FAILED或ROLLBACK_FAILED时，才能回退升级。  **默认取值**：不涉及
	Action UpgradeV3RequestBodyAction `json:"action"`

	// **参数解释**：四层规格ID。仅action为start时生效。  **约束限制**： - 负载均衡器有四层监听器时该字段必须指定。 - l4_flavor_id和l7_flavor_id不能同时为空。  **取值范围**：不涉及  **默认取值**：不涉及
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	// **参数解释**：七层规格ID。仅action为start时生效。  **约束限制**： - 负载均衡器有七层监听器时该字段必须指定。 - l4_flavor_id和l7_flavor_id不能同时为空。  **取值范围**：不涉及  **默认取值**：不涉及
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// **参数解释**：可用区列表。仅在action为start时生效，且action为start时，该字段必传。 可通过GET https://{ELB_Endpoint}/v3/{project_id}/elb/availability-zones 接口来查询可用区集合列表。创建负载均衡器时，从查询结果选择某一个可用区集合，并从中选择一个或多个可用区。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`

	// **参数解释**：双栈类型负载均衡器所在子网的IPv6网络ID。 可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的id得到。  **约束限制**： - 若实例升级到独享型后期望使用IPv6功能，则升级时该字段必传。 - ipv6_vip_virsubnet_id所属VPC必须与原共享型实例所属VPC相同。 - ipv6_vip_virsubnet_id所属子网需要开启IPv6。  **取值范围**：不涉及  **默认取值**：不涉及 [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// **参数解释**：负载均衡器的IPv6私网IP。  **约束限制**： - 该地址必须包含在所属子网的IPv6网段内，且未被占用。 - 传入ipv6_vip_address时必须传入ipv6_vip_virsubnet_id。 - 不传入ipv6_vip_address，但传入ipv6_vip_virsubnet_id，则自动分配IPv6私网IP。 - 不传入ipv6_vip_address，且不传ipv6_vip_virsubnet_id，则不分配私网IP，ipv6_vip_address=null。  **取值范围**：不涉及  **默认取值**：不涉及 [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	// **参数解释**：下联面子网的网络ID列表。 可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到。  **约束限制**： - 仅action为start时生效。 - 若不指定该字段，则选择vip_subnet_cidr_id子网对应的网络ID。 - 下联面子网必须属于该LB所在的VPC。  **取值范围**：不涉及  **默认取值**：不涉及
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`
}

func (o UpgradeV3RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeV3RequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeV3RequestBody", string(data)}, " ")
}

type UpgradeV3RequestBodyAction struct {
	value string
}

type UpgradeV3RequestBodyActionEnum struct {
	START    UpgradeV3RequestBodyAction
	COMPLETE UpgradeV3RequestBodyAction
	ROLLBACK UpgradeV3RequestBodyAction
}

func GetUpgradeV3RequestBodyActionEnum() UpgradeV3RequestBodyActionEnum {
	return UpgradeV3RequestBodyActionEnum{
		START: UpgradeV3RequestBodyAction{
			value: "start",
		},
		COMPLETE: UpgradeV3RequestBodyAction{
			value: "complete",
		},
		ROLLBACK: UpgradeV3RequestBodyAction{
			value: "rollback",
		},
	}
}

func (c UpgradeV3RequestBodyAction) Value() string {
	return c.value
}

func (c UpgradeV3RequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradeV3RequestBodyAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
