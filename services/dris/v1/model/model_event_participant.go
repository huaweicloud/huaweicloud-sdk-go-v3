package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventParticipant 交通参与者
type EventParticipant struct {

	// **参数说明**：交通参与者的具体类型描述。  **取值范围**：  - unknown：未知  - motor：机动车  - non-motor：非机动车  - pedestrian：行人
	PtcType *EventParticipantPtcType `json:"ptc_type,omitempty"`

	// **参数说明**：车牌号。
	PlateNo *string `json:"plate_no,omitempty"`

	// **参数说明**：对应车辆被检测到超速或者慢行时的速度小。单位为0.02米每秒。值为8191时代表无效数值。
	Speed *int32 `json:"speed,omitempty"`

	// **参数说明**：车辆类型。参考 [车辆基本类型参数说明](https://support.huaweicloud.com/api-v2x/v2x_04_0162.html)。
	VehicleClass *int32 `json:"vehicle_class,omitempty"`

	// **参数说明**：机动车车辆类型。参考[机动车车辆类型](https://support.huaweicloud.com/api-v2x/v2x_04_0179.html)。
	GatVehicleClass *string `json:"gat_vehicle_class,omitempty"`

	// **参数说明**：感知设备识别的id，具体表示为机动车轨迹ID。
	TrackId *int64 `json:"track_id,omitempty"`

	// **参数说明**：事件发生的所在车道
	LaneNo *int32 `json:"lane_no,omitempty"`

	// **参数说明**：目标检测框信息列表。
	TargetRects *[]TargetRect `json:"target_rects,omitempty"`
}

func (o EventParticipant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventParticipant struct{}"
	}

	return strings.Join([]string{"EventParticipant", string(data)}, " ")
}

type EventParticipantPtcType struct {
	value string
}

type EventParticipantPtcTypeEnum struct {
	UNKNOWN    EventParticipantPtcType
	MOTOR      EventParticipantPtcType
	NON_MOTOR  EventParticipantPtcType
	PEDESTRIAN EventParticipantPtcType
}

func GetEventParticipantPtcTypeEnum() EventParticipantPtcTypeEnum {
	return EventParticipantPtcTypeEnum{
		UNKNOWN: EventParticipantPtcType{
			value: "unknown",
		},
		MOTOR: EventParticipantPtcType{
			value: "motor",
		},
		NON_MOTOR: EventParticipantPtcType{
			value: "non-motor",
		},
		PEDESTRIAN: EventParticipantPtcType{
			value: "pedestrian",
		},
	}
}

func (c EventParticipantPtcType) Value() string {
	return c.value
}

func (c EventParticipantPtcType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventParticipantPtcType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
