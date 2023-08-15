package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KeystoneListEndpointsRequest Request Object
type KeystoneListEndpointsRequest struct {

	// 终端节点平面。可能取值为：public、internal或admin。public： 用户可在公共网络接口上看到。internal：用户可在内部网络接口上看到。admin：管理员可以在安全的网络接口上看到。
	Interface *KeystoneListEndpointsRequestInterface `json:"interface,omitempty"`

	// 服务ID。
	ServiceId *string `json:"service_id,omitempty"`
}

func (o KeystoneListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListEndpointsRequest", string(data)}, " ")
}

type KeystoneListEndpointsRequestInterface struct {
	value string
}

type KeystoneListEndpointsRequestInterfaceEnum struct {
	PUBLIC   KeystoneListEndpointsRequestInterface
	INTERNAL KeystoneListEndpointsRequestInterface
	ADMIN    KeystoneListEndpointsRequestInterface
}

func GetKeystoneListEndpointsRequestInterfaceEnum() KeystoneListEndpointsRequestInterfaceEnum {
	return KeystoneListEndpointsRequestInterfaceEnum{
		PUBLIC: KeystoneListEndpointsRequestInterface{
			value: "public",
		},
		INTERNAL: KeystoneListEndpointsRequestInterface{
			value: "internal",
		},
		ADMIN: KeystoneListEndpointsRequestInterface{
			value: "admin",
		},
	}
}

func (c KeystoneListEndpointsRequestInterface) Value() string {
	return c.value
}

func (c KeystoneListEndpointsRequestInterface) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeystoneListEndpointsRequestInterface) UnmarshalJSON(b []byte) error {
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
