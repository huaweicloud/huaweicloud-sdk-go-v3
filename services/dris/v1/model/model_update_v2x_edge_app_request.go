package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2xEdgeAppRequest Request Object
type UpdateV2xEdgeAppRequest struct {

	// **参数说明**：应用唯一ID，升级边缘应用前应先部署边缘应用，方法参见：[部署边缘应用](https://support.huaweicloud.com/api-v2x/v2x_04_0112.html)。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）、美元符号（$）的组合。
	EdgeAppId string `json:"edge_app_id"`

	// **参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和连接符（-）的组合，长度36。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：Edge ID，用于唯一标识一个Edge，创建Edge后获得。方法参见 [创建Edge](https://support.huaweicloud.com/api-v2x/v2x_04_0078.html)。
	V2xEdgeId string `json:"v2x_edge_id"`

	Body *ModifyV2XEdgeAppDto `json:"body,omitempty"`
}

func (o UpdateV2xEdgeAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2xEdgeAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateV2xEdgeAppRequest", string(data)}, " ")
}
