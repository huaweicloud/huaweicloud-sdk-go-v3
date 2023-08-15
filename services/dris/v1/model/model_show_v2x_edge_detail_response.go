package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2xEdgeDetailResponse Response Object
type ShowV2xEdgeDetailResponse struct {

	// **参数说明**：Edge ID，用于唯一标识一个Edge。
	V2xEdgeId *string `json:"v2x_edge_id,omitempty"`

	// **参数说明**：名称。  **取值范围**：长度不超过128，只允许中文、字母、数字、以及_.-等字符的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**：Edge描述。
	Description *string `json:"description,omitempty"`

	// **参数说明**：序列号。  **取值范围**：长度不超过64，只允许字母、数字、以及_等字符的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：网络I，例如127.0.0.1。
	Ip *string `json:"ip,omitempty"`

	// ITS800,ATLAS 端口号
	Port *int32 `json:"port,omitempty"`

	// **参数说明**：硬件类型。  **取值范围**：ITS800 或者 ATLAS。
	HardwareType *string `json:"hardware_type,omitempty"`

	// **参数说明**：安装位置编码，由用户自定义。  **取值范围**：长度不低于1不超过128，只允许字母、数字、下划线（_）的组合。
	PositionDescription *string `json:"position_description,omitempty"`

	Location *Location `json:"location,omitempty"`

	// **参数说明**：Edge关联的摄像头列表。
	Cameras *[]IdAndStatus `json:"cameras,omitempty"`

	// **参数说明**：Edge关联的雷达列表。
	Radars *[]IdAndStatus `json:"radars,omitempty"`

	// **参数说明**：Edge关联的本地RSU列表。
	LocalRsus *[]IdAndStatus `json:"local_rsus,omitempty"`

	EdgeGeneralConfig *EdgeGeneralConfigInResponse `json:"edge_general_config,omitempty"`

	// **参数说明**：Edge高级配置，Json格式
	EdgeAdvanceConfig *interface{} `json:"edge_advance_config,omitempty"`

	// **参数说明**：状态。  **取值范围**： - UNINSTALLED： 待部署 - INSTALLED：部署中 - OFFLINE：离线 - ONLINE：在线： - UPGRADING：升级中 - DELETING：删除中
	Status *string `json:"status,omitempty"`

	// **参数说明**：业务通道状态。  **取值范围**： - ONLINE：在线 - OFFLINE：离线 - INITIAL：初始化
	ChannelStatus *string `json:"channel_status,omitempty"`

	// **参数说明**：边缘管理服务返回的node_id，用于关联EdgeManager的资源。
	NodeId *string `json:"node_id,omitempty"`

	// **参数说明**：创建时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：创建时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o ShowV2xEdgeDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2xEdgeDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowV2xEdgeDetailResponse", string(data)}, " ")
}
