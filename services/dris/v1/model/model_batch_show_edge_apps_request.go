package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowEdgeAppsRequest Request Object
type BatchShowEdgeAppsRequest struct {

	// **参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和连接符（-）的组合，长度36。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：应用唯一ID。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）、美元符号（$）的组合。
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// **参数说明**：分页查询时的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：每页记录数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o BatchShowEdgeAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowEdgeAppsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowEdgeAppsRequest", string(data)}, " ")
}
