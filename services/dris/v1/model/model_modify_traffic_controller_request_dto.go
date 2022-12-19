package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyTrafficControllerRequestDto struct {

	// **参数说明**：名称。  **取值范围**：长度不超过128，只允许汉字、字母、数字、以及_-等字符的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**：描述
	Description *string `json:"description,omitempty"`

	// **参数说明**：定义纬度数值，北纬为正，南纬为负，单位°，精度小数点后7位。
	Lat float32 `json:"lat,omitempty"`

	// **参数说明**：定义经度数值。东经为正，西经为负，单位°，精度小数点后7位。
	Lon float32 `json:"lon,omitempty"`

	// **参数说明**：定义海拔高程，可选，单位为分米。
	Ele float32 `json:"ele,omitempty"`

	// **参数说明**：位置说明。  **取值范围**：长度不超过128，只允许字母、数字、以及_等字符的组合。
	PosDescription *string `json:"pos_description,omitempty"`

	// **参数说明**：架设方式。  **取值范围**： - columnar：柱式 - road-side-attach：路侧附着式 - cantilever：悬臂式 - gantry：门架式 - lane-above-attach：车行道上方附着式\"
	InstallationMode *string `json:"installation_mode,omitempty"`

	// **参数说明**：所属道路名称，比如高速名称。  **取值范围**：长度不超过64，只允许汉字、字母、数字、以及_-等字符的组合。
	RoadName *string `json:"road_name,omitempty"`

	// **参数说明**：信号机设备所属路段ID。  **取值范围**：长度等于30，只允许大写字母、数字。
	LinkId *string `json:"link_id,omitempty"`
}

func (o ModifyTrafficControllerRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTrafficControllerRequestDto struct{}"
	}

	return strings.Join([]string{"ModifyTrafficControllerRequestDto", string(data)}, " ")
}
