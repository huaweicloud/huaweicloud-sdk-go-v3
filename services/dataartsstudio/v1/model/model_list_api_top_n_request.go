package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApiTopNRequest Request Object
type ListApiTopNRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListApiTopNRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// api编号。
	ApiId string `json:"api_id"`

	// 集群编号。
	InstanceId *string `json:"instance_id,omitempty"`

	// 开始时间（13位时间戳）。
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）。
	EndTime int64 `json:"end_time"`

	// 时间单位。
	TimeUnit ListApiTopNRequestTimeUnit `json:"time_unit"`

	// 前几名。
	TopNum int32 `json:"top_num"`

	// 排序条件。
	OrderBy ListApiTopNRequestOrderBy `json:"order_by"`
}

func (o ListApiTopNRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiTopNRequest struct{}"
	}

	return strings.Join([]string{"ListApiTopNRequest", string(data)}, " ")
}

type ListApiTopNRequestDlmType struct {
	value string
}

type ListApiTopNRequestDlmTypeEnum struct {
	SHARED    ListApiTopNRequestDlmType
	EXCLUSIVE ListApiTopNRequestDlmType
}

func GetListApiTopNRequestDlmTypeEnum() ListApiTopNRequestDlmTypeEnum {
	return ListApiTopNRequestDlmTypeEnum{
		SHARED: ListApiTopNRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListApiTopNRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListApiTopNRequestDlmType) Value() string {
	return c.value
}

func (c ListApiTopNRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiTopNRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ListApiTopNRequestTimeUnit struct {
	value string
}

type ListApiTopNRequestTimeUnitEnum struct {
	HOUR ListApiTopNRequestTimeUnit
	DAY  ListApiTopNRequestTimeUnit
}

func GetListApiTopNRequestTimeUnitEnum() ListApiTopNRequestTimeUnitEnum {
	return ListApiTopNRequestTimeUnitEnum{
		HOUR: ListApiTopNRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ListApiTopNRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ListApiTopNRequestTimeUnit) Value() string {
	return c.value
}

func (c ListApiTopNRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiTopNRequestTimeUnit) UnmarshalJSON(b []byte) error {
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

type ListApiTopNRequestOrderBy struct {
	value string
}

type ListApiTopNRequestOrderByEnum struct {
	CALL_NUM              ListApiTopNRequestOrderBy
	SUCCESS_NUM           ListApiTopNRequestOrderBy
	FAIL_NUM              ListApiTopNRequestOrderBy
	LEGAL_NUM             ListApiTopNRequestOrderBy
	ILLEGAL_NUM           ListApiTopNRequestOrderBy
	COST_TIME_AVG         ListApiTopNRequestOrderBy
	SUCCESS_COST_TIME_AVG ListApiTopNRequestOrderBy
	FAIL_COST_TIME_AVG    ListApiTopNRequestOrderBy
	SUCCESS_RATE          ListApiTopNRequestOrderBy
	FAIL_RATE             ListApiTopNRequestOrderBy
	LEGAL_RATE            ListApiTopNRequestOrderBy
	ILLEGAL_RATE          ListApiTopNRequestOrderBy
}

func GetListApiTopNRequestOrderByEnum() ListApiTopNRequestOrderByEnum {
	return ListApiTopNRequestOrderByEnum{
		CALL_NUM: ListApiTopNRequestOrderBy{
			value: "CALL_NUM",
		},
		SUCCESS_NUM: ListApiTopNRequestOrderBy{
			value: "SUCCESS_NUM",
		},
		FAIL_NUM: ListApiTopNRequestOrderBy{
			value: "FAIL_NUM",
		},
		LEGAL_NUM: ListApiTopNRequestOrderBy{
			value: "LEGAL_NUM",
		},
		ILLEGAL_NUM: ListApiTopNRequestOrderBy{
			value: "ILLEGAL_NUM",
		},
		COST_TIME_AVG: ListApiTopNRequestOrderBy{
			value: "COST_TIME_AVG",
		},
		SUCCESS_COST_TIME_AVG: ListApiTopNRequestOrderBy{
			value: "SUCCESS_COST_TIME_AVG",
		},
		FAIL_COST_TIME_AVG: ListApiTopNRequestOrderBy{
			value: "FAIL_COST_TIME_AVG",
		},
		SUCCESS_RATE: ListApiTopNRequestOrderBy{
			value: "SUCCESS_RATE",
		},
		FAIL_RATE: ListApiTopNRequestOrderBy{
			value: "FAIL_RATE",
		},
		LEGAL_RATE: ListApiTopNRequestOrderBy{
			value: "LEGAL_RATE",
		},
		ILLEGAL_RATE: ListApiTopNRequestOrderBy{
			value: "ILLEGAL_RATE",
		},
	}
}

func (c ListApiTopNRequestOrderBy) Value() string {
	return c.value
}

func (c ListApiTopNRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiTopNRequestOrderBy) UnmarshalJSON(b []byte) error {
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
