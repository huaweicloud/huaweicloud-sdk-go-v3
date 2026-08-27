package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlgorithmsRequest Request Object
type ListAlgorithmsRequest struct {

	// **参数解释**： 排序规则，目前默认创建时间降序。 **约束限制**： 不涉及 **取值范围**： - DESC：降序 - ASC：升序 **默认取值**： DESC
	Order *string `json:"order,omitempty"`

	// **参数解释**： 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]。 **约束限制**： 不涉及 **取值范围**： [1,1000] **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]。 **约束限制**： 不涉及 **取值范围**： [0,100000000] **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 编程语言，可选python,c++,java **约束限制**： 不涉及 **取值范围**： python,c++,java **默认取值**： 0
	Lang *string `json:"lang,omitempty"`

	// **参数解释**： 算法id **约束限制**： 不涉及 **取值范围**： 长度[0,64] **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 算法名称 **约束限制**： 不涉及 **取值范围**： 长度[0,128] **默认取值**： 0
	Name *string `json:"name,omitempty"`

	// **参数解释**： 用户名 **约束限制**： 不涉及 **取值范围**： 长度[0,256] **默认取值**： 0
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 可见性 **约束限制**： 不涉及 **取值范围**： PUBLIC, PRIVATE **默认取值**： 无
	Visibility *ListAlgorithmsRequestVisibility `json:"visibility,omitempty"`

	// **参数解释**： 创建时间过滤条件，初始过滤时间 **约束限制**： 不涉及 **取值范围**： [0,9999999999999] **默认取值**： 无
	CreateTimeStart *int64 `json:"create_time_start,omitempty"`

	// **参数解释**： 创建时间过滤条件，终止过滤时间 **约束限制**： 不涉及 **取值范围**： [0,9999999999999] **默认取值**： 无
	CreateTimeEnd *int64 `json:"create_time_end,omitempty"`
}

func (o ListAlgorithmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlgorithmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlgorithmsRequest", string(data)}, " ")
}

type ListAlgorithmsRequestVisibility struct {
	value string
}

type ListAlgorithmsRequestVisibilityEnum struct {
	PUBLIC  ListAlgorithmsRequestVisibility
	PRIVATE ListAlgorithmsRequestVisibility
}

func GetListAlgorithmsRequestVisibilityEnum() ListAlgorithmsRequestVisibilityEnum {
	return ListAlgorithmsRequestVisibilityEnum{
		PUBLIC: ListAlgorithmsRequestVisibility{
			value: "PUBLIC",
		},
		PRIVATE: ListAlgorithmsRequestVisibility{
			value: "PRIVATE",
		},
	}
}

func (c ListAlgorithmsRequestVisibility) Value() string {
	return c.value
}

func (c ListAlgorithmsRequestVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlgorithmsRequestVisibility) UnmarshalJSON(b []byte) error {
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
