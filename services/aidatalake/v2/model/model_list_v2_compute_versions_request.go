package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListV2ComputeVersionsRequest Request Object
type ListV2ComputeVersionsRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：计算资源ID。可通过[查询计算资源列表]获取计算资源ID(ListV2Computes.xml)。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ComputeId *string `json:"compute_id,omitempty"`

	// **参数解释**：操作结果，不传则查询所有状态的操作记录。 **约束限制**：不涉及。 **取值范围**：   - SUCCESS：操作成功。   - EXECUTING：执行中。   - FAILED：操作失败。 **默认取值**：不涉及。
	OperationResult *ListV2ComputeVersionsRequestOperationResult `json:"operation_result,omitempty"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListV2ComputeVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2ComputeVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListV2ComputeVersionsRequest", string(data)}, " ")
}

type ListV2ComputeVersionsRequestOperationResult struct {
	value string
}

type ListV2ComputeVersionsRequestOperationResultEnum struct {
	SUCCESS   ListV2ComputeVersionsRequestOperationResult
	EXECUTING ListV2ComputeVersionsRequestOperationResult
	FAILED    ListV2ComputeVersionsRequestOperationResult
}

func GetListV2ComputeVersionsRequestOperationResultEnum() ListV2ComputeVersionsRequestOperationResultEnum {
	return ListV2ComputeVersionsRequestOperationResultEnum{
		SUCCESS: ListV2ComputeVersionsRequestOperationResult{
			value: "SUCCESS",
		},
		EXECUTING: ListV2ComputeVersionsRequestOperationResult{
			value: "EXECUTING",
		},
		FAILED: ListV2ComputeVersionsRequestOperationResult{
			value: "FAILED",
		},
	}
}

func (c ListV2ComputeVersionsRequestOperationResult) Value() string {
	return c.value
}

func (c ListV2ComputeVersionsRequestOperationResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListV2ComputeVersionsRequestOperationResult) UnmarshalJSON(b []byte) error {
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
