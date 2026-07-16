package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScaleUpHyperinstanceResponse Response Object
type ScaleUpHyperinstanceResponse struct {

	// **参数解释**：创建时间。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：超节点集群网络ID。 **约束限制**：不涉及。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **默认取值**：不涉及。
	HpsClusterId *string `json:"hps_cluster_id,omitempty"`

	// **参数解释**：超节点ID。 **约束限制**：不涉及。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **默认取值**：不涉及。
	HpsId *string `json:"hps_id,omitempty"`

	// **参数解释**：Lite Server超节点ID。 **约束限制**：不涉及。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：实例名称。 **约束限制**：不涉及。 **取值范围**：^[-_.a-zA-Z0-9]{1,64}$。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：订单ID。 **约束限制**：不涉及。 **取值范围**：^[a-zA-Z0-9]{1,64}$。 **默认取值**：不涉及。
	OrderId *string `json:"order_id,omitempty"`

	// **参数解释**：超节点实例状态。 **约束限制**：不涉及。 **取值范围**： - PROVISIONING：超节点的创建请求已被接受，但是仍在创建过程中； - ACTIVE：超节点处于活动状态，其资源可被使用； - ERROR：超节点创建失败； - REIMAGING：超节点切换操作系统中； - TERMINATING：资源释放中； - TERMINATED：超节点资源已经被释放，其资源不再可用。 **默认取值**：不涉及。
	Status *ScaleUpHyperinstanceResponseStatus `json:"status,omitempty"`

	// **参数解释**：超节点子节点实例列表。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Servers *[]ServerResponse `json:"servers,omitempty"`

	// **参数解释**：创建时间。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UpdateAt       *int64 `json:"update_at,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ScaleUpHyperinstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleUpHyperinstanceResponse struct{}"
	}

	return strings.Join([]string{"ScaleUpHyperinstanceResponse", string(data)}, " ")
}

type ScaleUpHyperinstanceResponseStatus struct {
	value string
}

type ScaleUpHyperinstanceResponseStatusEnum struct {
	PROVISIONING ScaleUpHyperinstanceResponseStatus
	ACTIVE       ScaleUpHyperinstanceResponseStatus
	ERROR        ScaleUpHyperinstanceResponseStatus
	REIMAGING    ScaleUpHyperinstanceResponseStatus
	TERMINATING  ScaleUpHyperinstanceResponseStatus
	TERMINATED   ScaleUpHyperinstanceResponseStatus
}

func GetScaleUpHyperinstanceResponseStatusEnum() ScaleUpHyperinstanceResponseStatusEnum {
	return ScaleUpHyperinstanceResponseStatusEnum{
		PROVISIONING: ScaleUpHyperinstanceResponseStatus{
			value: "PROVISIONING",
		},
		ACTIVE: ScaleUpHyperinstanceResponseStatus{
			value: "ACTIVE",
		},
		ERROR: ScaleUpHyperinstanceResponseStatus{
			value: "ERROR",
		},
		REIMAGING: ScaleUpHyperinstanceResponseStatus{
			value: "REIMAGING",
		},
		TERMINATING: ScaleUpHyperinstanceResponseStatus{
			value: "TERMINATING",
		},
		TERMINATED: ScaleUpHyperinstanceResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c ScaleUpHyperinstanceResponseStatus) Value() string {
	return c.value
}

func (c ScaleUpHyperinstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScaleUpHyperinstanceResponseStatus) UnmarshalJSON(b []byte) error {
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
