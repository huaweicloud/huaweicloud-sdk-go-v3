package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetAvailableZoneResponseBody **参数解释**： 可用区详情
type GetAvailableZoneResponseBody struct {

	// **参数解释**： 可用区ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 可用区名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 可用区按所在区域显示的名称 **取值范围**： 不涉及
	DisplayName *string `json:"displayName,omitempty"`

	// **参数解释**： 可用区所属的可用区组，一个可用区可能属于多个可用区组 **取值范围**： 不涉及
	AzGroupIds *[]string `json:"azGroupIds,omitempty"`

	// **参数解释**： EIP所属的组，IES边缘场景为可用区ID，中心区统一为“center” **取值范围**： 不涉及
	PublicBorderGroup *string `json:"PublicBorderGroup,omitempty"`

	// **参数解释**： 可用区分类 **取值范围**： - Default: 中心云可用区 - IES: 专属云可用区 - HomeZone: 智能边缘云可用区
	Category *GetAvailableZoneResponseBodyCategory `json:"category,omitempty"`

	// **参数解释**： 可用区名称别名 **取值范围**： 不涉及
	Alias *string `json:"alias,omitempty"`
}

func (o GetAvailableZoneResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAvailableZoneResponseBody struct{}"
	}

	return strings.Join([]string{"GetAvailableZoneResponseBody", string(data)}, " ")
}

type GetAvailableZoneResponseBodyCategory struct {
	value string
}

type GetAvailableZoneResponseBodyCategoryEnum struct {
	DEFAULT   GetAvailableZoneResponseBodyCategory
	IES       GetAvailableZoneResponseBodyCategory
	HOME_ZONE GetAvailableZoneResponseBodyCategory
}

func GetGetAvailableZoneResponseBodyCategoryEnum() GetAvailableZoneResponseBodyCategoryEnum {
	return GetAvailableZoneResponseBodyCategoryEnum{
		DEFAULT: GetAvailableZoneResponseBodyCategory{
			value: "Default",
		},
		IES: GetAvailableZoneResponseBodyCategory{
			value: "IES",
		},
		HOME_ZONE: GetAvailableZoneResponseBodyCategory{
			value: "HomeZone",
		},
	}
}

func (c GetAvailableZoneResponseBodyCategory) Value() string {
	return c.value
}

func (c GetAvailableZoneResponseBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetAvailableZoneResponseBodyCategory) UnmarshalJSON(b []byte) error {
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
