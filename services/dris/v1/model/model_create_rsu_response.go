package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRsuResponse Response Object
type CreateRsuResponse struct {

	// **参数说明**：RSU的唯一标识符，在平台创建RSU时由平台生成。
	RsuId *string `json:"rsu_id,omitempty"`

	// **参数说明**：RSU的名字。  **取值范围**：长度不低于1不超过128，只允许中文、字母、数字、下划线（_）、连接符（-）的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**：RSU的描述。  **取值范围**：只允许中文、字母、数字、下划线（_）、中文分号（；）、中文冒号（：）、中文问号（？）、中文感叹号（！）中文逗号（，）、中文句号（。）、英文分号（;）、英文冒号（:）、英文逗号（,）、英文句号（.）、英文问号（?）、英文感叹号（!）、顿号（、）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：RSU的设备序列号。  **取值范围**：只允许字母、数字、下划线（_）的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：最后修改的时间。  格式：yyyy-MM-dd'T'HH:mm:ss'Z'  例如 2020-09-01T01:37:01Z
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`

	// **参数说明**：创建的时间。  格式：yyyy-MM-dd'T'HH:mm:ss'Z'  例如 2020-09-01T01:37:01Z
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：最后的在线时间。  格式：yyyy-MM-dd'T'HH:mm:ss'Z'  例如 2020-09-01T01:37:01Z
	LastOnlineTime *sdktime.SdkTime `json:"last_online_time,omitempty"`

	// **参数说明**：RSU的IP。满足IP的格式，例如127.0.0.1。
	Ip *string `json:"ip,omitempty"`

	// **参数说明**：安装位置编码，由用户自定义。  **取值范围**：长度不低于1不超过128，只允许字母、数字、下划线（_）的组合。
	PositionDescription *string `json:"position_description,omitempty"`

	Location *RsuLocation `json:"location,omitempty"`

	// **参数说明**：RSU设备状态。  **取值范围**：  - ONLINE：在线  - OFFLINE：离线  - INITIAL：初始化  - UNKNOWN：未知
	Status *string `json:"status,omitempty"`

	// **参数说明**：RSU型号ID，用于唯一标识一个RSU型号，在平台创建RSU型号后由平台分配获得，获取方法可参见 [创建RSU型号](https://support.huaweicloud.com/api-v2x/v2x_04_0020.html)。  **取值范围**：长度不低于1不超过36，只允许字母、数字、连接符（-）的组合。  **该字段仅供使用MQTT协议RSU设备的用户输入。使用websocket协议RSU设备的用户需忽略此字段。**
	RsuModelId *string `json:"rsu_model_id,omitempty"`

	// **参数说明**：在地图中，rsu所在区域对应的路口ID，也即区域ID拼接路口ID，格式为：region-node_id。其中路网最基本的构成即节点和节点之间连接的路段。节点可以是路口，也可以是一条 路的端点。一个节点的ID在同一个区域内是唯一的。
	IntersectionId *string `json:"intersection_id,omitempty"`

	// **参数说明**：RSU可关联的Edge的数量。
	RelatedEdgeNum *int32 `json:"related_edge_num,omitempty"`

	// **参数说明**：RSU的软件版本，由RSU上报其软件版本。
	SoftwareVersion *string `json:"software_version,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateRsuResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRsuResponse struct{}"
	}

	return strings.Join([]string{"CreateRsuResponse", string(data)}, " ")
}
