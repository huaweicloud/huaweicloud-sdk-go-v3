package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateBackupResourcePackageResponse Response Object
type CreateBackupResourcePackageResponse struct {

	// **参数解释**：  创建备份资源包订单ID。  **取值范围**：  不涉及。
	OrderId *string `json:"order_id,omitempty"`

	// **参数解释**：  备份资源包规格码。  **取值范围**：  不涉及。
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**：  备份资源包数量。  **取值范围**：  1-10。
	Num *int32 `json:"num,omitempty"`

	// **参数解释**：  订购周期类型。  **取值范围**：  - month：包月。 - year：包年。
	PeriodType *CreateBackupResourcePackageResponsePeriodType `json:"period_type,omitempty"`

	// **参数解释**：  订购时间长度。  **取值范围**：  - \"period_type\"为\"month\"时，取值为1~9。 - \"period_type\"为\"year\"时，取值为1~3。
	PeriodNum      *int32 `json:"period_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateBackupResourcePackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupResourcePackageResponse struct{}"
	}

	return strings.Join([]string{"CreateBackupResourcePackageResponse", string(data)}, " ")
}

type CreateBackupResourcePackageResponsePeriodType struct {
	value string
}

type CreateBackupResourcePackageResponsePeriodTypeEnum struct {
	MONTH CreateBackupResourcePackageResponsePeriodType
	YEAR  CreateBackupResourcePackageResponsePeriodType
}

func GetCreateBackupResourcePackageResponsePeriodTypeEnum() CreateBackupResourcePackageResponsePeriodTypeEnum {
	return CreateBackupResourcePackageResponsePeriodTypeEnum{
		MONTH: CreateBackupResourcePackageResponsePeriodType{
			value: "month",
		},
		YEAR: CreateBackupResourcePackageResponsePeriodType{
			value: "year",
		},
	}
}

func (c CreateBackupResourcePackageResponsePeriodType) Value() string {
	return c.value
}

func (c CreateBackupResourcePackageResponsePeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBackupResourcePackageResponsePeriodType) UnmarshalJSON(b []byte) error {
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
