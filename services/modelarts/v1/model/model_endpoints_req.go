package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointsReq 本地IDE（如PyCharm、VS Code）或SSH客户端，通过SSH远程接入Notebook实例时需要的相关配置。
type EndpointsReq struct {

	// **参数解释**：支持的服务。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下：  - NOTEBOOK：可以通过https协议访问Notebook。  - SSH：可以通过SSH协议远程连接Notebook。  **默认取值**：不涉及。
	Service *string `json:"service,omitempty"`

	// **参数解释**：SSH密钥对名称，可以在云服务器控制台（ECS）“密钥对”页面创建和查看。 **约束限制**：不涉及。
	KeyPairNames *[]string `json:"key_pair_names,omitempty"`
}

func (o EndpointsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointsReq struct{}"
	}

	return strings.Join([]string{"EndpointsReq", string(data)}, " ")
}
