package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ColumnInfo 列信息
type ColumnInfo struct {

	// 列名
	ColumnName *string `json:"column_name,omitempty"`

	// 列类型
	ColumnType *string `json:"column_type,omitempty"`

	// 列的键类型。
	KeyType *string `json:"key_type,omitempty"`

	// 列映射后的名称
	ColumnMappedName *string `json:"column_mapped_name,omitempty"`

	// 该列是否过滤
	Status *bool `json:"status,omitempty"`

	// 该列是否partitionKey
	Partition *bool `json:"partition,omitempty"`

	// 列的脱敏算法类型（主键和唯一键不支持）。取值： - NAME ：姓名中间遮盖。 - PHONE ：电话中间遮盖。 - IDENTITY_CARD ：证件号码中间遮盖。  - BANK_CARD ：银行卡号中间遮盖。 - EMAIL ：邮箱前缀首字母后遮盖。 - ADDRESS ：地址后缀遮盖。  - DATE ：仅保留年代日期替换为1月1日。  - NUMERIC ：数字变0。 - SECRET_CONTENT ：大段文本全替换。
	DesensitizationAlgorithm *ColumnInfoDesensitizationAlgorithm `json:"desensitization_algorithm,omitempty"`
}

func (o ColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnInfo struct{}"
	}

	return strings.Join([]string{"ColumnInfo", string(data)}, " ")
}

type ColumnInfoDesensitizationAlgorithm struct {
	value string
}

type ColumnInfoDesensitizationAlgorithmEnum struct {
	NAME           ColumnInfoDesensitizationAlgorithm
	PHONE          ColumnInfoDesensitizationAlgorithm
	IDENTITY_CARD  ColumnInfoDesensitizationAlgorithm
	BANK_CARD      ColumnInfoDesensitizationAlgorithm
	EMAIL          ColumnInfoDesensitizationAlgorithm
	ADDRESS        ColumnInfoDesensitizationAlgorithm
	DATE           ColumnInfoDesensitizationAlgorithm
	NUMERIC        ColumnInfoDesensitizationAlgorithm
	SECRET_CONTENT ColumnInfoDesensitizationAlgorithm
}

func GetColumnInfoDesensitizationAlgorithmEnum() ColumnInfoDesensitizationAlgorithmEnum {
	return ColumnInfoDesensitizationAlgorithmEnum{
		NAME: ColumnInfoDesensitizationAlgorithm{
			value: "name",
		},
		PHONE: ColumnInfoDesensitizationAlgorithm{
			value: "phone",
		},
		IDENTITY_CARD: ColumnInfoDesensitizationAlgorithm{
			value: "identity_card",
		},
		BANK_CARD: ColumnInfoDesensitizationAlgorithm{
			value: "bank_card",
		},
		EMAIL: ColumnInfoDesensitizationAlgorithm{
			value: "email",
		},
		ADDRESS: ColumnInfoDesensitizationAlgorithm{
			value: "address",
		},
		DATE: ColumnInfoDesensitizationAlgorithm{
			value: "date",
		},
		NUMERIC: ColumnInfoDesensitizationAlgorithm{
			value: "numeric",
		},
		SECRET_CONTENT: ColumnInfoDesensitizationAlgorithm{
			value: "secret_content",
		},
	}
}

func (c ColumnInfoDesensitizationAlgorithm) Value() string {
	return c.value
}

func (c ColumnInfoDesensitizationAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ColumnInfoDesensitizationAlgorithm) UnmarshalJSON(b []byte) error {
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
