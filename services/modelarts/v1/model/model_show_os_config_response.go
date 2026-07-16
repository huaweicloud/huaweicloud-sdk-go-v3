package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOsConfigResponse Response Object
type ShowOsConfigResponse struct {

	// **参数解释**：网络配置项。
	NetworkCidrs *[]string `json:"networkCidrs,omitempty"`

	// **参数解释**：用户可创建网络个数配额。 **取值范围**：不涉及
	NetworkQuota *int32 `json:"networkQuota,omitempty"`

	// **参数解释**：用户可创建资源池个数配额。 **取值范围**：不涉及
	PoolQuota *int32 `json:"poolQuota,omitempty"`

	// **参数解释**：当前环境/局点是否支持创建高可用资源池。 **取值范围**： - true：支持 - false：不支持
	PoolHighAvailable *bool `json:"poolHighAvailable,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o ShowOsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOsConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowOsConfigResponse", string(data)}, " ")
}
