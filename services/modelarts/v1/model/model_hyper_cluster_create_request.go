package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HyperClusterCreateRequest struct {

	// **参数解释**：hyper cluster的名称。 **取值范围**：^[-_.a-zA-Z0-9]{1,64}$。
	Name string `json:"name"`

	// **参数解释**：hyper cluster的ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	HyperClusterSubnetId *string `json:"hyper_cluster_subnet_id,omitempty"`

	// **参数解释**：服务器类型。 **约束限制**：不涉及。 **取值范围**： - HPS：超节点服务 - ECS：弹性云服务 **默认取值**：不涉及。
	Type *HyperClusterCreateRequestType `json:"type,omitempty"`
}

func (o HyperClusterCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperClusterCreateRequest struct{}"
	}

	return strings.Join([]string{"HyperClusterCreateRequest", string(data)}, " ")
}

type HyperClusterCreateRequestType struct {
	value string
}

type HyperClusterCreateRequestTypeEnum struct {
	HPS HyperClusterCreateRequestType
	ECS HyperClusterCreateRequestType
}

func GetHyperClusterCreateRequestTypeEnum() HyperClusterCreateRequestTypeEnum {
	return HyperClusterCreateRequestTypeEnum{
		HPS: HyperClusterCreateRequestType{
			value: "HPS",
		},
		ECS: HyperClusterCreateRequestType{
			value: "ECS",
		},
	}
}

func (c HyperClusterCreateRequestType) Value() string {
	return c.value
}

func (c HyperClusterCreateRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HyperClusterCreateRequestType) UnmarshalJSON(b []byte) error {
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
