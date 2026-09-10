package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AttachDevServerPortResponse Response Object
type AttachDevServerPortResponse struct {

	// **参数解释**：端口MAC地址，由系统分配。
	MacAddr *string `json:"mac_addr,omitempty"`

	// **参数解释**：网卡ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	PortId *string `json:"port_id,omitempty"`

	// **参数解释**：端口状态。 **取值范围**： - ACTIVE：端口处于活动状态，可以正常进行网络通信。 - BUILD：端口正在创建或配置中。 - DOWN：端口处于非活动状态，不能进行网络通信。
	PortState *AttachDevServerPortResponsePortState `json:"port_state,omitempty"`

	// **参数解释**：端口所在子网ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	VirsubnetId    *string `json:"virsubnet_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachDevServerPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDevServerPortResponse struct{}"
	}

	return strings.Join([]string{"AttachDevServerPortResponse", string(data)}, " ")
}

type AttachDevServerPortResponsePortState struct {
	value string
}

type AttachDevServerPortResponsePortStateEnum struct {
	ACTIVE AttachDevServerPortResponsePortState
	BUILD  AttachDevServerPortResponsePortState
	DOWN   AttachDevServerPortResponsePortState
}

func GetAttachDevServerPortResponsePortStateEnum() AttachDevServerPortResponsePortStateEnum {
	return AttachDevServerPortResponsePortStateEnum{
		ACTIVE: AttachDevServerPortResponsePortState{
			value: "ACTIVE",
		},
		BUILD: AttachDevServerPortResponsePortState{
			value: "BUILD",
		},
		DOWN: AttachDevServerPortResponsePortState{
			value: "DOWN",
		},
	}
}

func (c AttachDevServerPortResponsePortState) Value() string {
	return c.value
}

func (c AttachDevServerPortResponsePortState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachDevServerPortResponsePortState) UnmarshalJSON(b []byte) error {
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
