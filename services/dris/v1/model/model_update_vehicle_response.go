package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVehicleResponse Response Object
type UpdateVehicleResponse struct {

	// **参数说明**：车辆唯一标识符。  **取值范围**：长度不超过128，只允许字母、数字、以及_-等字符的组合。
	VehicleId *string `json:"vehicle_id,omitempty"`

	// **参数说明**：车牌号。  **取值范围**：长度最小1最大64，支持中文、字母、数字、下划线（_）、横杠（-）的组合。
	PlateNo *string `json:"plate_no,omitempty"`

	// **参数说明**：VIN码，车辆的17位VIN码。  **取值范围**：长度不超过17，只允许字母、数字字符的组合。
	Vin *string `json:"vin,omitempty"`

	// \"**参数说明**：车载OBU的唯一标识。  **取值范围**：长度不超过128，只允许字母、数字、以及_-等字符的组合。
	ObuId *string `json:"obu_id,omitempty"`

	// **参数说明**：IMEI，OBU上电子序列号。  **取值范围**：长度最小1最大255，支持纯数字的组合。
	Imei *string `json:"imei,omitempty"`

	// **参数说明**：车俩品牌。  **取值范围**：长度最小1最大64，支持中文、字母、数字、下划线（_）、横杠（-）的组合。
	Brand *string `json:"brand,omitempty"`

	// **参数说明**：车牌型号。  **取值范围**：长度最小1最大64，支持字母、数字、下划线（_）、横杠（-）的组合。
	Model *string `json:"model,omitempty"`

	// **参数说明**：车辆年款。  **取值范围**：长度最小1最大64，支持纯数字的组合。
	Style *string `json:"style,omitempty"`

	// **参数说明**：定义车辆的燃料动力类。  **取值范围**： - unknownFuel：未知 - gasoline：汽油 - ethanol：乙醇 - diesel：柴油 - electric：电动 - hybrid：混合燃料类型 - hydrogen：氢气 - natGasLiquid：液化天然气 - natGasComp：压缩天然气 - propane：丙烷\"
	FuelType *string `json:"fuel_type,omitempty"`

	// **参数说明**：车辆颜色。  **取值范围**： - black：黑色 - white：白色 - gray：灰色 - red：红色 - blue：蓝色 - yellow：黄色 - orange：橙色 - brown：棕色 - green：绿色 - purple：紫色 - cyan：青色 - pink：粉红色 - transparent：透明色 - other：其他\"
	Color *string `json:"color,omitempty"`

	// **参数说明**：车辆颜色。  **取值范围**： - black：黑色 - white：白色 - blue：蓝色 - yellow：黄色 - green：绿色
	PlateColor *string `json:"plate_color,omitempty"`

	// **参数说明**：车辆接入网络的方式。  **取值范围**： - 5g - 4g  - 3g  - 2g - pc5Only  - pc5And5g  - pc5And4g  - pc5And3g  - pc5And2g
	AccessType *string `json:"access_type,omitempty"`

	// **参数说明**：描述。  **取值范围**：长度不超过2048，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：最后修改的时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`

	// **参数说明**：创建的时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：最后的在线时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	LastOnlineTime *sdktime.SdkTime `json:"last_online_time,omitempty"`

	// **参数说明**：设备状态。  **取值范围**： - ONLINE：在线 - OFFLINE：离线 - UNKNOWN：未知状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVehicleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVehicleResponse struct{}"
	}

	return strings.Join([]string{"UpdateVehicleResponse", string(data)}, " ")
}
