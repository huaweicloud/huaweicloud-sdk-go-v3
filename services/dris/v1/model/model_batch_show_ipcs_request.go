package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchShowIpcsRequest Request Object
type BatchShowIpcsRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：分页查询时的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：分页查询时每页显示的记录数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：摄像机的状态。  **取值范围**：  - ONLINE：在线   - OFFLINE：离线  - INITIAL：初始化  - UNKNOWN：未知   - SLEEP：休眠
	Status *BatchShowIpcsRequestStatus `json:"status,omitempty"`

	// **参数说明**：Edge ID，用于唯一标识一个Edge，创建Edge后获得。方法参见 [创建Edge](https://support.huaweicloud.com/api-v2x/v2x_04_0078.html)。
	V2xEdgeId *string `json:"v2x_edge_id,omitempty"`
}

func (o BatchShowIpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowIpcsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowIpcsRequest", string(data)}, " ")
}

type BatchShowIpcsRequestStatus struct {
	value string
}

type BatchShowIpcsRequestStatusEnum struct {
	ONLINE  BatchShowIpcsRequestStatus
	OFFLINE BatchShowIpcsRequestStatus
	INITIAL BatchShowIpcsRequestStatus
	UNKNOWN BatchShowIpcsRequestStatus
	SLEEP   BatchShowIpcsRequestStatus
}

func GetBatchShowIpcsRequestStatusEnum() BatchShowIpcsRequestStatusEnum {
	return BatchShowIpcsRequestStatusEnum{
		ONLINE: BatchShowIpcsRequestStatus{
			value: "ONLINE",
		},
		OFFLINE: BatchShowIpcsRequestStatus{
			value: "OFFLINE",
		},
		INITIAL: BatchShowIpcsRequestStatus{
			value: "INITIAL",
		},
		UNKNOWN: BatchShowIpcsRequestStatus{
			value: "UNKNOWN",
		},
		SLEEP: BatchShowIpcsRequestStatus{
			value: "SLEEP",
		},
	}
}

func (c BatchShowIpcsRequestStatus) Value() string {
	return c.value
}

func (c BatchShowIpcsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchShowIpcsRequestStatus) UnmarshalJSON(b []byte) error {
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
