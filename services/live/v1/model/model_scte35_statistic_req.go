package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Scte35StatisticReq 请求频道scte35信号的字段。
type Scte35StatisticReq struct {

	// 信号类型：all/splice_insert/time_signal。
	Type Scte35StatisticReqType `json:"type"`

	// 查询信号的起始时间，unix time，单位：秒。
	StartTime int64 `json:"start_time"`

	// 查询信号的结束时间，unix time，单位：秒；实际使用使用比start_time大。
	EndTime int64 `json:"end_time"`
}

func (o Scte35StatisticReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scte35StatisticReq struct{}"
	}

	return strings.Join([]string{"Scte35StatisticReq", string(data)}, " ")
}

type Scte35StatisticReqType struct {
	value string
}

type Scte35StatisticReqTypeEnum struct {
	ALL           Scte35StatisticReqType
	SPLICE_INSERT Scte35StatisticReqType
	TIME_SIGNAL   Scte35StatisticReqType
}

func GetScte35StatisticReqTypeEnum() Scte35StatisticReqTypeEnum {
	return Scte35StatisticReqTypeEnum{
		ALL: Scte35StatisticReqType{
			value: "all",
		},
		SPLICE_INSERT: Scte35StatisticReqType{
			value: "splice_insert",
		},
		TIME_SIGNAL: Scte35StatisticReqType{
			value: "time_signal",
		},
	}
}

func (c Scte35StatisticReqType) Value() string {
	return c.value
}

func (c Scte35StatisticReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Scte35StatisticReqType) UnmarshalJSON(b []byte) error {
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
