package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceStorageInfoResponse Response Object
type ShowInstanceStorageInfoResponse struct {

	// **参数解释**：  已使用存储空间大小，单位是GB。  **取值范围**：  大于等于0。
	Used *float64 `json:"used,omitempty"`

	// **参数解释**：  包年/包月购买的存储空间大小，单位是GB。  **取值范围**：  大于0。
	PrepaidSize *float64 `json:"prepaid_size,omitempty"`

	// **参数解释**：  实例可使用的存储空间上限，单位是GB。  **取值范围**：  大于0。
	Limit *float64 `json:"limit,omitempty"`

	// **参数解释**：  存储类型。  **取值范围**：  - POOL：DL6存储类型。 - STANDARDPOOL：DL5存储类型。
	Type           *ShowInstanceStorageInfoResponseType `json:"type,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowInstanceStorageInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStorageInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceStorageInfoResponse", string(data)}, " ")
}

type ShowInstanceStorageInfoResponseType struct {
	value string
}

type ShowInstanceStorageInfoResponseTypeEnum struct {
	POOL         ShowInstanceStorageInfoResponseType
	STANDARDPOOL ShowInstanceStorageInfoResponseType
}

func GetShowInstanceStorageInfoResponseTypeEnum() ShowInstanceStorageInfoResponseTypeEnum {
	return ShowInstanceStorageInfoResponseTypeEnum{
		POOL: ShowInstanceStorageInfoResponseType{
			value: "POOL",
		},
		STANDARDPOOL: ShowInstanceStorageInfoResponseType{
			value: "STANDARDPOOL",
		},
	}
}

func (c ShowInstanceStorageInfoResponseType) Value() string {
	return c.value
}

func (c ShowInstanceStorageInfoResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceStorageInfoResponseType) UnmarshalJSON(b []byte) error {
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
