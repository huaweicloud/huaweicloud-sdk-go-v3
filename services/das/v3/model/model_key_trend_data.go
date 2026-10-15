package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type KeyTrendData struct {

	// key名称
	Name *string `json:"name,omitempty"`

	// 统计维度。size/storage/qps/traffic
	StatType *KeyTrendDataStatType `json:"stat_type,omitempty"`

	// 时间粒度（毫秒）
	Interval *int64 `json:"interval,omitempty"`

	// 趋势数据点列表
	TrendList *[]TrendPoint `json:"trend_list,omitempty"`
}

func (o KeyTrendData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyTrendData struct{}"
	}

	return strings.Join([]string{"KeyTrendData", string(data)}, " ")
}

type KeyTrendDataStatType struct {
	value string
}

type KeyTrendDataStatTypeEnum struct {
	SIZE    KeyTrendDataStatType
	STORAGE KeyTrendDataStatType
	QPS     KeyTrendDataStatType
	TRAFFIC KeyTrendDataStatType
}

func GetKeyTrendDataStatTypeEnum() KeyTrendDataStatTypeEnum {
	return KeyTrendDataStatTypeEnum{
		SIZE: KeyTrendDataStatType{
			value: "size",
		},
		STORAGE: KeyTrendDataStatType{
			value: "storage",
		},
		QPS: KeyTrendDataStatType{
			value: "qps",
		},
		TRAFFIC: KeyTrendDataStatType{
			value: "traffic",
		},
	}
}

func (c KeyTrendDataStatType) Value() string {
	return c.value
}

func (c KeyTrendDataStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeyTrendDataStatType) UnmarshalJSON(b []byte) error {
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
