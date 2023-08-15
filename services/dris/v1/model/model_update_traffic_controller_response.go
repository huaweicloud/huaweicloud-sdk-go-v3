package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrafficControllerResponse Response Object
type UpdateTrafficControllerResponse struct {

	// **参数说明**：信号机设备ID，全局唯一。
	TrafficControllerId *string `json:"traffic_controller_id,omitempty"`

	// \"**参数说明**：序列号。  **取值范围**：长度不超过64，只允许字母、数字、以及_等字符的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：名称。
	Name *string `json:"name,omitempty"`

	// **参数说明**：描述。  **取值范围**：长度不超过2048，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：定义纬度数值，北纬为正，南纬为负，单位°，精度小数点后7位。
	Lat float32 `json:"lat,omitempty"`

	// **参数说明**：定义经度数值。东经为正，西经为负，单位°，精度小数点后7位。
	Lon float32 `json:"lon,omitempty"`

	// **参数说明**：定义海拔高程，可选，单位为分米。
	Ele float32 `json:"ele,omitempty"`

	// **参数说明**：位置说明。  **取值范围**：长度不超过128，只允许字母、数字、以及_等字符的组合。
	PosDescription *string `json:"pos_description,omitempty"`

	// **参数说明**：架设方式。  **取值范围**： - columnar：柱式 - road-side-attach：路侧附着式 - cantilever：悬臂式 - gantry：门架式 - lane-above-attach：车行道上方附着式
	InstallationMode *string `json:"installation_mode,omitempty"`

	// **参数说明**：所属道路名称，比如高速名称。  **取值范围**：长度不超过64，只允许汉字、字母、数字、以及_-等字符的组合。
	RoadName *string `json:"road_name,omitempty"`

	// **参数说明**：信号机设备所属路段ID。  **取值范围**：长度等于30，只允许大写字母、数字。
	LinkId *string `json:"link_id,omitempty"`

	// **参数说明**：设备状态。  **取值范围**： - ONLINE：在线 - OFFLINE：离线 - INITIAL：初始化
	Status *string `json:"status,omitempty"`

	// **参数说明**：最后修改的时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`

	// **参数说明**：创建的时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：最后的在线时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	LastOnlineTime *sdktime.SdkTime `json:"last_online_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateTrafficControllerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrafficControllerResponse struct{}"
	}

	return strings.Join([]string{"UpdateTrafficControllerResponse", string(data)}, " ")
}
