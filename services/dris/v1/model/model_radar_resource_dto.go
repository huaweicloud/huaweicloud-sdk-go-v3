package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type RadarResourceDto struct {

	// **参数说明**：雷达ID
	RadarId *string `json:"radar_id,omitempty"`

	// **参数说明**：名称  **取值范围**：长度不小于1，不大于128的汉字、英文字母、数字、下划线（_）和横杠（-）的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**：所属的EdgeId  **取值范围**：长度不小于1，不大于128的英文字母、数字、下划线（_）和横杠（-）的组合。
	V2xEdgeId *string `json:"v2x_edge_id,omitempty"`

	// **参数说明**：设备IP
	Ip *string `json:"ip,omitempty"`

	// **参数说明**：状态值  **取值范围**：   - ONLINE：在线   - OFFLINE：离线   - UNKNOWN：未知   - SLEEP：睡眠
	Status *RadarResourceDtoStatus `json:"status,omitempty"`

	// **参数说明**：设备编号  **取值范围**：长度不小于1，不大于64的英文字母、数字和下划线（_）的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：位置描述  **取值范围**：长度不小于1，不大于128的英文字母、数字和下划线（_）的组合。
	PositionDescription *string `json:"position_description,omitempty"`

	// **参数说明**：最后修改的时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：最后修改时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`

	// **参数说明**：最后的在线时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z' 例如 2020-09-01T01:37:01Z
	LastOnlineTime *sdktime.SdkTime `json:"last_online_time,omitempty"`
}

func (o RadarResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RadarResourceDto struct{}"
	}

	return strings.Join([]string{"RadarResourceDto", string(data)}, " ")
}

type RadarResourceDtoStatus struct {
	value string
}

type RadarResourceDtoStatusEnum struct {
	ONLINE  RadarResourceDtoStatus
	OFFLINE RadarResourceDtoStatus
	UNKNOWN RadarResourceDtoStatus
	SLEEP   RadarResourceDtoStatus
}

func GetRadarResourceDtoStatusEnum() RadarResourceDtoStatusEnum {
	return RadarResourceDtoStatusEnum{
		ONLINE: RadarResourceDtoStatus{
			value: "ONLINE",
		},
		OFFLINE: RadarResourceDtoStatus{
			value: "OFFLINE",
		},
		UNKNOWN: RadarResourceDtoStatus{
			value: "UNKNOWN",
		},
		SLEEP: RadarResourceDtoStatus{
			value: "SLEEP",
		},
	}
}

func (c RadarResourceDtoStatus) Value() string {
	return c.value
}

func (c RadarResourceDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RadarResourceDtoStatus) UnmarshalJSON(b []byte) error {
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
