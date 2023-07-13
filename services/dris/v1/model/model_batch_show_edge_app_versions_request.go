package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchShowEdgeAppVersionsRequest Request Object
type BatchShowEdgeAppVersionsRequest struct {

	// **参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和连接符（-）的组合，长度36。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：用户自定义应用唯一ID。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）、美元符号（$）的组合。
	EdgeAppId string `json:"edge_app_id"`

	// **参数说明**：应用版本搜索关键字。
	Version *string `json:"version,omitempty"`

	// **参数说明**：分页查询时的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：每页记录数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：应用版本状态。  **取值范围**：  - DRAFT：草稿  - PUBLISHED：发布  - OFF_SHELF：下线
	State *BatchShowEdgeAppVersionsRequestState `json:"state,omitempty"`
}

func (o BatchShowEdgeAppVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowEdgeAppVersionsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowEdgeAppVersionsRequest", string(data)}, " ")
}

type BatchShowEdgeAppVersionsRequestState struct {
	value string
}

type BatchShowEdgeAppVersionsRequestStateEnum struct {
	DRAFT     BatchShowEdgeAppVersionsRequestState
	PUBLISHED BatchShowEdgeAppVersionsRequestState
	OFF_SHELF BatchShowEdgeAppVersionsRequestState
}

func GetBatchShowEdgeAppVersionsRequestStateEnum() BatchShowEdgeAppVersionsRequestStateEnum {
	return BatchShowEdgeAppVersionsRequestStateEnum{
		DRAFT: BatchShowEdgeAppVersionsRequestState{
			value: "DRAFT",
		},
		PUBLISHED: BatchShowEdgeAppVersionsRequestState{
			value: "PUBLISHED",
		},
		OFF_SHELF: BatchShowEdgeAppVersionsRequestState{
			value: "OFF_SHELF",
		},
	}
}

func (c BatchShowEdgeAppVersionsRequestState) Value() string {
	return c.value
}

func (c BatchShowEdgeAppVersionsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchShowEdgeAppVersionsRequestState) UnmarshalJSON(b []byte) error {
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
