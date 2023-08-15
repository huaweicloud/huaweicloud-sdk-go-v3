package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2xEdgeResponse Response Object
type UpdateV2xEdgeResponse struct {

	// **参数说明**：Edge ID，用于唯一标识一个Edge
	V2xEdgeId *string `json:"v2x_edge_id,omitempty"`

	// **参数说明**：名称。  **取值范围**：长度不低于1不超过128，只允许中文、字母、数字、下划线（_）、连接符（-）的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**：Edge描述。  **取值范围**：长度不超过255，只允许中文、字母、数字、下划线（_）、中文分号（；）、中文冒号（：）、中文问号（？）、中文感叹号（！）中文逗号（，）、中文句号（。）、英文引号（;）、英文冒号（:）、英文逗号（,）、英文句号（.）、英文问号（?）、英文感叹号（!）、顿号（、）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：设备编码。  **取值范围**：长度不超过64，只允许字母、数字、以及_等字符的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：网络IP，例如127.0.0.1。
	Ip *string `json:"ip,omitempty"`

	// ITS800,ATLAS 端口号
	Port *int32 `json:"port,omitempty"`

	// **参数说明**：硬件类型。  **取值范围**：ITS800 或者 ATLAS
	HardwareType *string `json:"hardware_type,omitempty"`

	// **参数说明**：安装位置编码，由用户自定义。  **取值范围**：长度不低于1不超过128，只允许字母、数字、下划线（_）的组合。
	PositionDescription *string `json:"position_description,omitempty"`

	Location *Location `json:"location,omitempty"`

	// **参数说明**：摄像头ID列表。
	CameraIds *[]string `json:"camera_ids,omitempty"`

	// **参数说明**：雷达ID列表。
	RadarIds *[]string `json:"radar_ids,omitempty"`

	// **参数说明**：Edge关联的本地RSU列表。
	LocalRsus *[]string `json:"local_rsus,omitempty"`

	EdgeGeneralConfig *EdgeGeneralConfigInResponse `json:"edge_general_config,omitempty"`

	// Edge高级配置，Json格式
	EdgeAdvanceConfig *interface{} `json:"edge_advance_config,omitempty"`

	// \"**参数说明**：状态。  **取值范围**： - UNINSTALLED： 待部署 - INSTALLED：部署中 - OFFLINE：离线 - ONLINE：在线： - UPGRADING：升级中 - DELETING：删除中
	Status *string `json:"status,omitempty"`

	// 边缘管理服务返回的node_id，用于关联EdgeManager的资源
	NodeId *string `json:"node_id,omitempty"`

	// **参数说明**：创建时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：创建时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o UpdateV2xEdgeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2xEdgeResponse struct{}"
	}

	return strings.Join([]string{"UpdateV2xEdgeResponse", string(data)}, " ")
}
