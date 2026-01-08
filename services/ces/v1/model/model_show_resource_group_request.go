package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourceGroupRequest Request Object
type ShowResourceGroupRequest struct {

	// **参数解释** 资源分组ID。 **约束限制** 不涉及 **取值范围** 以\"rg\"开头，后面跟着22个字母或数字 **默认取值** 不涉及
	GroupId string `json:"group_id"`

	// **参数解释** 资源分组健康状态 **约束限制** 不涉及 **取值范围** - health: 表示无告警 - unhealth: 表示告警中 - no_alarm_rule: 表示未设置告警规则 **默认取值** 不涉及
	Status *ShowResourceGroupRequestStatus `json:"status,omitempty"`

	// **参数解释** 资源类型，即命名空间，如弹性云服务器的资源命名空间为：SYS.ECS；各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。 **约束限制** 不涉及 **取值范围** 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度在 [3,32]个字符之间 **默认取值** 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释** 资源维度，如：弹性云服务器，则维度为instance_id，各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **约束限制** 不涉及 **取值范围** 包含字母、数字、_、-、/、#、\\或括号，长度为[1,131]个字符 **默认取值** 不涉及
	Dname *string `json:"dname,omitempty"`

	// **参数解释** 分页起始值 **约束限制** 不涉及 **取值范围** [0,9999999] **默认取值** 0
	Start *string `json:"start,omitempty"`

	// **参数解释** 单次查询的条数限制 **约束限制** 不涉及 **取值范围** [1,100] **默认取值** 100
	Limit *string `json:"limit,omitempty"`
}

func (o ShowResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceGroupRequest", string(data)}, " ")
}

type ShowResourceGroupRequestStatus struct {
	value string
}

type ShowResourceGroupRequestStatusEnum struct {
	HEALTH        ShowResourceGroupRequestStatus
	UNHEALTH      ShowResourceGroupRequestStatus
	NO_ALARM_RULE ShowResourceGroupRequestStatus
}

func GetShowResourceGroupRequestStatusEnum() ShowResourceGroupRequestStatusEnum {
	return ShowResourceGroupRequestStatusEnum{
		HEALTH: ShowResourceGroupRequestStatus{
			value: "health",
		},
		UNHEALTH: ShowResourceGroupRequestStatus{
			value: "unhealth",
		},
		NO_ALARM_RULE: ShowResourceGroupRequestStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ShowResourceGroupRequestStatus) Value() string {
	return c.value
}

func (c ShowResourceGroupRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceGroupRequestStatus) UnmarshalJSON(b []byte) error {
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
