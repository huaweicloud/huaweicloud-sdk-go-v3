package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAppsTopRequest Request Object
type ListAppsTopRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ListAppsTopRequestDlmType `json:"Dlm-Type"`

	// 开始时间（13位时间戳）
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）
	EndTime int64 `json:"end_time"`

	// 时间单位
	TimeUnit ListAppsTopRequestTimeUnit `json:"time_unit"`

	// 前几名
	TopNum int32 `json:"top_num"`

	// 排序条件
	OrderBy ListAppsTopRequestOrderBy `json:"order_by"`
}

func (o ListAppsTopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsTopRequest struct{}"
	}

	return strings.Join([]string{"ListAppsTopRequest", string(data)}, " ")
}

type ListAppsTopRequestDlmType struct {
	value string
}

type ListAppsTopRequestDlmTypeEnum struct {
	SHARED    ListAppsTopRequestDlmType
	EXCLUSIVE ListAppsTopRequestDlmType
}

func GetListAppsTopRequestDlmTypeEnum() ListAppsTopRequestDlmTypeEnum {
	return ListAppsTopRequestDlmTypeEnum{
		SHARED: ListAppsTopRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListAppsTopRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListAppsTopRequestDlmType) Value() string {
	return c.value
}

func (c ListAppsTopRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAppsTopRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ListAppsTopRequestTimeUnit struct {
	value string
}

type ListAppsTopRequestTimeUnitEnum struct {
	HOUR ListAppsTopRequestTimeUnit
	DAY  ListAppsTopRequestTimeUnit
}

func GetListAppsTopRequestTimeUnitEnum() ListAppsTopRequestTimeUnitEnum {
	return ListAppsTopRequestTimeUnitEnum{
		HOUR: ListAppsTopRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ListAppsTopRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ListAppsTopRequestTimeUnit) Value() string {
	return c.value
}

func (c ListAppsTopRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAppsTopRequestTimeUnit) UnmarshalJSON(b []byte) error {
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

type ListAppsTopRequestOrderBy struct {
	value string
}

type ListAppsTopRequestOrderByEnum struct {
	CALL_NUM              ListAppsTopRequestOrderBy
	SUCCESS_NUM           ListAppsTopRequestOrderBy
	FAIL_NUM              ListAppsTopRequestOrderBy
	LEGAL_NUM             ListAppsTopRequestOrderBy
	ILLEGAL_NUM           ListAppsTopRequestOrderBy
	COST_TIME_AVG         ListAppsTopRequestOrderBy
	SUCCESS_COST_TIME_AVG ListAppsTopRequestOrderBy
	FAIL_COST_TIME_AVG    ListAppsTopRequestOrderBy
	SUCCESS_RATE          ListAppsTopRequestOrderBy
	FAIL_RATE             ListAppsTopRequestOrderBy
	LEGAL_RATE            ListAppsTopRequestOrderBy
	ILLEGAL_RATE          ListAppsTopRequestOrderBy
}

func GetListAppsTopRequestOrderByEnum() ListAppsTopRequestOrderByEnum {
	return ListAppsTopRequestOrderByEnum{
		CALL_NUM: ListAppsTopRequestOrderBy{
			value: "CALL_NUM",
		},
		SUCCESS_NUM: ListAppsTopRequestOrderBy{
			value: "SUCCESS_NUM",
		},
		FAIL_NUM: ListAppsTopRequestOrderBy{
			value: "FAIL_NUM",
		},
		LEGAL_NUM: ListAppsTopRequestOrderBy{
			value: "LEGAL_NUM",
		},
		ILLEGAL_NUM: ListAppsTopRequestOrderBy{
			value: "ILLEGAL_NUM",
		},
		COST_TIME_AVG: ListAppsTopRequestOrderBy{
			value: "COST_TIME_AVG",
		},
		SUCCESS_COST_TIME_AVG: ListAppsTopRequestOrderBy{
			value: "SUCCESS_COST_TIME_AVG",
		},
		FAIL_COST_TIME_AVG: ListAppsTopRequestOrderBy{
			value: "FAIL_COST_TIME_AVG",
		},
		SUCCESS_RATE: ListAppsTopRequestOrderBy{
			value: "SUCCESS_RATE",
		},
		FAIL_RATE: ListAppsTopRequestOrderBy{
			value: "FAIL_RATE",
		},
		LEGAL_RATE: ListAppsTopRequestOrderBy{
			value: "LEGAL_RATE",
		},
		ILLEGAL_RATE: ListAppsTopRequestOrderBy{
			value: "ILLEGAL_RATE",
		},
	}
}

func (c ListAppsTopRequestOrderBy) Value() string {
	return c.value
}

func (c ListAppsTopRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAppsTopRequestOrderBy) UnmarshalJSON(b []byte) error {
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
