package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// V2XEdge模型资源列表返回对象
type V2XEdgeListResponseDto struct {

	// **参数说明**：Edge ID，用于唯一标识一个Edge。
	V2xEdgeId *string `json:"v2x_edge_id,omitempty"`

	// **参数说明**：V2XEdge的名称  **取值范围**：长度不超过128，只允许中文、字母、数字、以及_.-等字符的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**：设备编码。  **取值范围**：长度不超过64，只允许字母、数字、以及_等字符的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：网络IP，例如127.0.0.1。
	Ip *string `json:"ip,omitempty"`

	// **参数说明**：安装位置编码，由用户自定义。  **取值范围**：长度不低于1不超过128，只允许字母、数字、下划线（_）的组合。
	PositionDescription *string `json:"position_description,omitempty"`

	Location *Location `json:"location,omitempty"`

	// **参数说明**：状态。  **取值范围**： - UNINSTALLED： 待部署 - INSTALLED：部署中 - OFFLINE：离线 - ONLINE：在线： - UPGRADING：升级中 - DELETING：删除中
	Status *string `json:"status,omitempty"`

	// **参数说明**：业务通道状态。
	ChannelStatus *V2XEdgeListResponseDtoChannelStatus `json:"channel_status,omitempty"`

	// **参数说明**：创建时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`
}

func (o V2XEdgeListResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2XEdgeListResponseDto struct{}"
	}

	return strings.Join([]string{"V2XEdgeListResponseDto", string(data)}, " ")
}

type V2XEdgeListResponseDtoChannelStatus struct {
	value string
}

type V2XEdgeListResponseDtoChannelStatusEnum struct {
	ONLINE  V2XEdgeListResponseDtoChannelStatus
	OFFLINE V2XEdgeListResponseDtoChannelStatus
	UNKNOWN V2XEdgeListResponseDtoChannelStatus
}

func GetV2XEdgeListResponseDtoChannelStatusEnum() V2XEdgeListResponseDtoChannelStatusEnum {
	return V2XEdgeListResponseDtoChannelStatusEnum{
		ONLINE: V2XEdgeListResponseDtoChannelStatus{
			value: "ONLINE",
		},
		OFFLINE: V2XEdgeListResponseDtoChannelStatus{
			value: "OFFLINE",
		},
		UNKNOWN: V2XEdgeListResponseDtoChannelStatus{
			value: "UNKNOWN",
		},
	}
}

func (c V2XEdgeListResponseDtoChannelStatus) Value() string {
	return c.value
}

func (c V2XEdgeListResponseDtoChannelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2XEdgeListResponseDtoChannelStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
