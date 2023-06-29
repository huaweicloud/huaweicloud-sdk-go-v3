package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2xEdgeAppResponse Response Object
type UpdateV2xEdgeAppResponse struct {

	// **参数说明**：Edge ID，用于唯一标识一个Edge。
	V2xEdgeId *string `json:"v2x_edge_id,omitempty"`

	// **参数说明**：用户自定义应用唯一ID。
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// **参数说明**：应用版本，比如1.0.0。
	AppVersion *string `json:"app_version,omitempty"`

	// **参数说明**：应用部署状态。  **取值范围**：  - UNINSTALLED：待部署  - INSTALLED：部署中  - OFFLINE：离线  - ONLINE：在线  - UPGRADING：升级中  - DELETING：删除中  - RUNNING：运行中
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateV2xEdgeAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2xEdgeAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateV2xEdgeAppResponse", string(data)}, " ")
}
