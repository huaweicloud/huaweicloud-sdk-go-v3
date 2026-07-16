package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CloudServer struct {

	// **参数解释**：服务器资源id，或超节点子节点id。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	Id *string `json:"id,omitempty"`

	// **参数解释**：Lite Server服务器类型。 **取值范围**： - BMS：裸金属服务器 - ECS：弹性云服务器 - HPS：超节点服务器
	Type *CloudServerType `json:"type,omitempty"`

	// **参数解释**：服务器所属的超节点资源id。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	HpsId *string `json:"hps_id,omitempty"`

	// **参数解释**：超节点子节点对应服务器资源id。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	HpsEcsId *string `json:"hps_ecs_id,omitempty"`
}

func (o CloudServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudServer struct{}"
	}

	return strings.Join([]string{"CloudServer", string(data)}, " ")
}

type CloudServerType struct {
	value string
}

type CloudServerTypeEnum struct {
	BMS CloudServerType
	ECS CloudServerType
	HPS CloudServerType
}

func GetCloudServerTypeEnum() CloudServerTypeEnum {
	return CloudServerTypeEnum{
		BMS: CloudServerType{
			value: "BMS",
		},
		ECS: CloudServerType{
			value: "ECS",
		},
		HPS: CloudServerType{
			value: "HPS",
		},
	}
}

func (c CloudServerType) Value() string {
	return c.value
}

func (c CloudServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudServerType) UnmarshalJSON(b []byte) error {
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
