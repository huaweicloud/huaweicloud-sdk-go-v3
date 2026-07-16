package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EndpointsRes 本地IDE（如PyCharm、VS Code）或SSH客户端，通过SSH远程接入Notebook实例时需要的相关配置。
type EndpointsRes struct {

	// **参数解释**：访问Notebook的途径。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：可以通过https协议访问Notebook。 - SSH：可以通过SSH协议远程连接Notebook。
	DevService *string `json:"dev_service,omitempty"`

	// **参数解释**：访问Notebook的途径。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：可以通过https协议访问Notebook。 - SSH：可以通过SSH协议远程连接Notebook。
	Service *EndpointsResService `json:"service,omitempty"`

	// **参数解释**：实例私有IP地址。 **取值范围**：不涉及。
	Uri *string `json:"uri,omitempty"`

	// **参数解释**：SSH密钥对名称列表，允许设置多个密钥对实现同时对SSH实例的访问。 **取值范围**：不涉及。
	KeyPairNames *[]string `json:"key_pair_names,omitempty"`
}

func (o EndpointsRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointsRes struct{}"
	}

	return strings.Join([]string{"EndpointsRes", string(data)}, " ")
}

type EndpointsResService struct {
	value string
}

type EndpointsResServiceEnum struct {
	NOTEBOOK EndpointsResService
	SSH      EndpointsResService
}

func GetEndpointsResServiceEnum() EndpointsResServiceEnum {
	return EndpointsResServiceEnum{
		NOTEBOOK: EndpointsResService{
			value: "NOTEBOOK",
		},
		SSH: EndpointsResService{
			value: "SSH",
		},
	}
}

func (c EndpointsResService) Value() string {
	return c.value
}

func (c EndpointsResService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointsResService) UnmarshalJSON(b []byte) error {
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
