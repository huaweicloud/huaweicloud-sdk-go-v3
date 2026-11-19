package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplicaBinlogReadConfigResponse Response Object
type ShowReplicaBinlogReadConfigResponse struct {

	// **参数解释**：  实例是否支持开启只读Binlog拉取。  **取值范围**：  - true：支持。 - false：不支持。
	Supported *bool `json:"supported,omitempty"`

	// **参数解释**：  只读Binlog拉取开关状态。  **取值范围**：  - true：开启。 - false：关闭。
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowReplicaBinlogReadConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicaBinlogReadConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowReplicaBinlogReadConfigResponse", string(data)}, " ")
}
