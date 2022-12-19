package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteV2XEdgeAppByEdgeAppIdRequest struct {

	// **参数说明**：用户自定义应用唯一ID。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）、美元符号（$）的组合。
	EdgeAppId string `json:"edge_app_id"`

	// **参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和连接符（-）的组合，长度36。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：Edge ID，用于唯一标识一个Edge，创建Edge后获得。方法参见 [创建Edge](https://support.huaweicloud.com/api-v2x/v2x_04_0078.html)。
	V2xEdgeId string `json:"v2x_edge_id"`
}

func (o DeleteV2XEdgeAppByEdgeAppIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2XEdgeAppByEdgeAppIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteV2XEdgeAppByEdgeAppIdRequest", string(data)}, " ")
}
