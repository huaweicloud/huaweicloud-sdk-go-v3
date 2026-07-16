package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Endpoints 本地IDE（如PyCharm、VSCode）或SSH客户端，通过SSH远程接入Notebook实例时需要的相关配置。
type Endpoints struct {

	// **参数解释**：支持的服务。 **取值范围**： - NOTEBOOK：可以通过https协议访问Notebook - SSH：可以通过SSH协议远程连接Notebook
	DevService *EndpointsDevService `json:"dev_service,omitempty"`

	// **参数解释**：通过应用专属URL直接打开应用进入远程开发模式。包含应用的各种扩展配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Extensions map[string]string `json:"extensions,omitempty"`

	// **参数解释**：SSH密钥对名称列表。允许设置多个密钥对实现同时对SSH实例的访问。 **约束限制**：不涉及。 **取值范围**：0 - 1024个密钥对 **默认取值**：不涉及。
	SshKeys *[]string `json:"ssh_keys,omitempty"`
}

func (o Endpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Endpoints struct{}"
	}

	return strings.Join([]string{"Endpoints", string(data)}, " ")
}

type EndpointsDevService struct {
	value string
}

type EndpointsDevServiceEnum struct {
	AI_FLOW      EndpointsDevService
	MA_STUDIO    EndpointsDevService
	NOTEBOOK     EndpointsDevService
	SSH          EndpointsDevService
	TENSOR_BOARD EndpointsDevService
	WEB_IDE      EndpointsDevService
}

func GetEndpointsDevServiceEnum() EndpointsDevServiceEnum {
	return EndpointsDevServiceEnum{
		AI_FLOW: EndpointsDevService{
			value: "AI_FLOW",
		},
		MA_STUDIO: EndpointsDevService{
			value: "MA_STUDIO",
		},
		NOTEBOOK: EndpointsDevService{
			value: "NOTEBOOK",
		},
		SSH: EndpointsDevService{
			value: "SSH",
		},
		TENSOR_BOARD: EndpointsDevService{
			value: "TENSOR_BOARD",
		},
		WEB_IDE: EndpointsDevService{
			value: "WEB_IDE",
		},
	}
}

func (c EndpointsDevService) Value() string {
	return c.value
}

func (c EndpointsDevService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointsDevService) UnmarshalJSON(b []byte) error {
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
