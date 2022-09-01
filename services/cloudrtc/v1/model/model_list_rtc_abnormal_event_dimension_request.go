package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListRtcAbnormalEventDimensionRequest struct {

	// 应用ID
	App string `json:"app" xml:"app"`

	// 房间ID
	RoomId *string `json:"room_id,omitempty" xml:"room_id"`

	// 分组类型，支持同时指定两种类型 - abnormal_type：异常类型 - abnormal_factor：异常因素
	Dimension *ListRtcAbnormalEventDimensionRequestDimension `json:"dimension,omitempty" xml:"dimension"`

	// 查询起始时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T06:00:00Z，不填写则默认读取过去1小时数据数据。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 查询结束时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T07:00:00Z，不填写则默认为当前时间。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o ListRtcAbnormalEventDimensionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcAbnormalEventDimensionRequest struct{}"
	}

	return strings.Join([]string{"ListRtcAbnormalEventDimensionRequest", string(data)}, " ")
}

type ListRtcAbnormalEventDimensionRequestDimension struct {
	value string
}

type ListRtcAbnormalEventDimensionRequestDimensionEnum struct {
	ABNORMAL_TYPE   ListRtcAbnormalEventDimensionRequestDimension
	ABNORMAL_FACTOR ListRtcAbnormalEventDimensionRequestDimension
}

func GetListRtcAbnormalEventDimensionRequestDimensionEnum() ListRtcAbnormalEventDimensionRequestDimensionEnum {
	return ListRtcAbnormalEventDimensionRequestDimensionEnum{
		ABNORMAL_TYPE: ListRtcAbnormalEventDimensionRequestDimension{
			value: "abnormal_type",
		},
		ABNORMAL_FACTOR: ListRtcAbnormalEventDimensionRequestDimension{
			value: "abnormal_factor",
		},
	}
}

func (c ListRtcAbnormalEventDimensionRequestDimension) Value() string {
	return c.value
}

func (c ListRtcAbnormalEventDimensionRequestDimension) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRtcAbnormalEventDimensionRequestDimension) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
