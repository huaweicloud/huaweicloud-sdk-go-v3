package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAiAppsRequest Request Object
type ListAiAppsRequest struct {

	// **参数解释：** 最大返回条目数 **约束限制：** 不涉及 **取值范围：** 1-100 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 跳过的条目数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 按应用名称过滤 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 按防护状态过滤 **约束限制：** 不涉及 **取值范围：** 0(禁用), 1(启用) **默认取值：** 不涉及
	ProtectStatus *ListAiAppsRequestProtectStatus `json:"protect_status,omitempty"`

	// **参数解释：** 按策略ID过滤 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId *string `json:"policy_id,omitempty"`
}

func (o ListAiAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAiAppsRequest", string(data)}, " ")
}

type ListAiAppsRequestProtectStatus struct {
	value int32
}

type ListAiAppsRequestProtectStatusEnum struct {
	E_0 ListAiAppsRequestProtectStatus
	E_1 ListAiAppsRequestProtectStatus
}

func GetListAiAppsRequestProtectStatusEnum() ListAiAppsRequestProtectStatusEnum {
	return ListAiAppsRequestProtectStatusEnum{
		E_0: ListAiAppsRequestProtectStatus{
			value: 0,
		}, E_1: ListAiAppsRequestProtectStatus{
			value: 1,
		},
	}
}

func (c ListAiAppsRequestProtectStatus) Value() int32 {
	return c.value
}

func (c ListAiAppsRequestProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAiAppsRequestProtectStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
