package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type IpcResponseDto struct {

	// **参数说明**：摄像头ID，console界面查询摄像头IPC列表中的设备Id。
	CameraId *string `json:"camera_id,omitempty"`

	// **参数说明**：Edge ID，用于唯一标识一个Edge，创建Edge后获得。方法参见 [创建Edge](https://support.huaweicloud.com/api-v2x/v2x_04_0078.html)。
	V2xEdgeId *string `json:"v2x_edge_id,omitempty"`

	// **参数说明**：摄像头名称。
	Name *string `json:"name,omitempty"`

	// **参数说明**：摄像头所感知的路口或者路段的Id。
	CrossId *string `json:"cross_id,omitempty"`

	// **参数说明**：摄像头聚焦类型。  - long：长焦  - short：短焦
	FocalType *string `json:"focal_type,omitempty"`

	// **参数说明**：摄像头连接的ITS800的互联编码。
	ParentConnectCode *string `json:"parent_connect_code,omitempty"`

	// **参数说明**：摄像头的互联编码。
	ConnectCode *string `json:"connect_code,omitempty"`

	// **参数说明**：描述。
	Description *string `json:"description,omitempty"`

	// **参数说明**：IPC的设备编码。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：该摄像头的ip地址。
	Ip *string `json:"ip,omitempty"`

	// **参数说明**：摄像机的状态。  **取值范围**：  - ONLINE：在线   - OFFLINE：离线  - INITIAL：初始化  - UNKNOWN：未知   - SLEEP：休眠
	Status *IpcResponseDtoStatus `json:"status,omitempty"`

	// **参数说明**：创建时间。            格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：最后修改时间。            格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`

	// **参数说明**：最后在线时间。          格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	LastOnlineTime *sdktime.SdkTime `json:"last_online_time,omitempty"`
}

func (o IpcResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpcResponseDto struct{}"
	}

	return strings.Join([]string{"IpcResponseDto", string(data)}, " ")
}

type IpcResponseDtoStatus struct {
	value string
}

type IpcResponseDtoStatusEnum struct {
	ONLINE  IpcResponseDtoStatus
	OFFLINE IpcResponseDtoStatus
	INITIAL IpcResponseDtoStatus
	UNKNOWN IpcResponseDtoStatus
	SLEEP   IpcResponseDtoStatus
}

func GetIpcResponseDtoStatusEnum() IpcResponseDtoStatusEnum {
	return IpcResponseDtoStatusEnum{
		ONLINE: IpcResponseDtoStatus{
			value: "ONLINE",
		},
		OFFLINE: IpcResponseDtoStatus{
			value: "OFFLINE",
		},
		INITIAL: IpcResponseDtoStatus{
			value: "INITIAL",
		},
		UNKNOWN: IpcResponseDtoStatus{
			value: "UNKNOWN",
		},
		SLEEP: IpcResponseDtoStatus{
			value: "SLEEP",
		},
	}
}

func (c IpcResponseDtoStatus) Value() string {
	return c.value
}

func (c IpcResponseDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpcResponseDtoStatus) UnmarshalJSON(b []byte) error {
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
