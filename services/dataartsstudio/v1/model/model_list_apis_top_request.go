package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApisTopRequest Request Object
type ListApisTopRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListApisTopRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// 集群编号。
	InstanceId *string `json:"instance_id,omitempty"`

	// 开始时间（13位时间戳）。
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）。
	EndTime int64 `json:"end_time"`

	// 时间单位。
	TimeUnit ListApisTopRequestTimeUnit `json:"time_unit"`

	// 前几名。
	TopNum int32 `json:"top_num"`

	// 排序条件。
	OrderBy ListApisTopRequestOrderBy `json:"order_by"`
}

func (o ListApisTopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisTopRequest struct{}"
	}

	return strings.Join([]string{"ListApisTopRequest", string(data)}, " ")
}

type ListApisTopRequestDlmType struct {
	value string
}

type ListApisTopRequestDlmTypeEnum struct {
	SHARED    ListApisTopRequestDlmType
	EXCLUSIVE ListApisTopRequestDlmType
}

func GetListApisTopRequestDlmTypeEnum() ListApisTopRequestDlmTypeEnum {
	return ListApisTopRequestDlmTypeEnum{
		SHARED: ListApisTopRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListApisTopRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListApisTopRequestDlmType) Value() string {
	return c.value
}

func (c ListApisTopRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisTopRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ListApisTopRequestTimeUnit struct {
	value string
}

type ListApisTopRequestTimeUnitEnum struct {
	HOUR ListApisTopRequestTimeUnit
	DAY  ListApisTopRequestTimeUnit
}

func GetListApisTopRequestTimeUnitEnum() ListApisTopRequestTimeUnitEnum {
	return ListApisTopRequestTimeUnitEnum{
		HOUR: ListApisTopRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ListApisTopRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ListApisTopRequestTimeUnit) Value() string {
	return c.value
}

func (c ListApisTopRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisTopRequestTimeUnit) UnmarshalJSON(b []byte) error {
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

type ListApisTopRequestOrderBy struct {
	value string
}

type ListApisTopRequestOrderByEnum struct {
	CALL_NUM              ListApisTopRequestOrderBy
	SUCCESS_NUM           ListApisTopRequestOrderBy
	FAIL_NUM              ListApisTopRequestOrderBy
	LEGAL_NUM             ListApisTopRequestOrderBy
	ILLEGAL_NUM           ListApisTopRequestOrderBy
	COST_TIME_AVG         ListApisTopRequestOrderBy
	SUCCESS_COST_TIME_AVG ListApisTopRequestOrderBy
	FAIL_COST_TIME_AVG    ListApisTopRequestOrderBy
	SUCCESS_RATE          ListApisTopRequestOrderBy
	FAIL_RATE             ListApisTopRequestOrderBy
	LEGAL_RATE            ListApisTopRequestOrderBy
	ILLEGAL_RATE          ListApisTopRequestOrderBy
}

func GetListApisTopRequestOrderByEnum() ListApisTopRequestOrderByEnum {
	return ListApisTopRequestOrderByEnum{
		CALL_NUM: ListApisTopRequestOrderBy{
			value: "CALL_NUM",
		},
		SUCCESS_NUM: ListApisTopRequestOrderBy{
			value: "SUCCESS_NUM",
		},
		FAIL_NUM: ListApisTopRequestOrderBy{
			value: "FAIL_NUM",
		},
		LEGAL_NUM: ListApisTopRequestOrderBy{
			value: "LEGAL_NUM",
		},
		ILLEGAL_NUM: ListApisTopRequestOrderBy{
			value: "ILLEGAL_NUM",
		},
		COST_TIME_AVG: ListApisTopRequestOrderBy{
			value: "COST_TIME_AVG",
		},
		SUCCESS_COST_TIME_AVG: ListApisTopRequestOrderBy{
			value: "SUCCESS_COST_TIME_AVG",
		},
		FAIL_COST_TIME_AVG: ListApisTopRequestOrderBy{
			value: "FAIL_COST_TIME_AVG",
		},
		SUCCESS_RATE: ListApisTopRequestOrderBy{
			value: "SUCCESS_RATE",
		},
		FAIL_RATE: ListApisTopRequestOrderBy{
			value: "FAIL_RATE",
		},
		LEGAL_RATE: ListApisTopRequestOrderBy{
			value: "LEGAL_RATE",
		},
		ILLEGAL_RATE: ListApisTopRequestOrderBy{
			value: "ILLEGAL_RATE",
		},
	}
}

func (c ListApisTopRequestOrderBy) Value() string {
	return c.value
}

func (c ListApisTopRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisTopRequestOrderBy) UnmarshalJSON(b []byte) error {
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
