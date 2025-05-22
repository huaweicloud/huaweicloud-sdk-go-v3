package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoDescriptions struct {

	// 自定义gop时长，单位秒，有效取值[0,10]. 默认值0，表示跟随源流，不支持7和9，值无效时按默认值处理
	GopDurationSeconds *VideoDescriptionsGopDurationSeconds `json:"gop_duration_seconds,omitempty"`
}

func (o VideoDescriptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoDescriptions struct{}"
	}

	return strings.Join([]string{"VideoDescriptions", string(data)}, " ")
}

type VideoDescriptionsGopDurationSeconds struct {
	value int32
}

type VideoDescriptionsGopDurationSecondsEnum struct {
	E_0  VideoDescriptionsGopDurationSeconds
	E_1  VideoDescriptionsGopDurationSeconds
	E_2  VideoDescriptionsGopDurationSeconds
	E_3  VideoDescriptionsGopDurationSeconds
	E_4  VideoDescriptionsGopDurationSeconds
	E_5  VideoDescriptionsGopDurationSeconds
	E_6  VideoDescriptionsGopDurationSeconds
	E_8  VideoDescriptionsGopDurationSeconds
	E_10 VideoDescriptionsGopDurationSeconds
}

func GetVideoDescriptionsGopDurationSecondsEnum() VideoDescriptionsGopDurationSecondsEnum {
	return VideoDescriptionsGopDurationSecondsEnum{
		E_0: VideoDescriptionsGopDurationSeconds{
			value: 0,
		}, E_1: VideoDescriptionsGopDurationSeconds{
			value: 1,
		}, E_2: VideoDescriptionsGopDurationSeconds{
			value: 2,
		}, E_3: VideoDescriptionsGopDurationSeconds{
			value: 3,
		}, E_4: VideoDescriptionsGopDurationSeconds{
			value: 4,
		}, E_5: VideoDescriptionsGopDurationSeconds{
			value: 5,
		}, E_6: VideoDescriptionsGopDurationSeconds{
			value: 6,
		}, E_8: VideoDescriptionsGopDurationSeconds{
			value: 8,
		}, E_10: VideoDescriptionsGopDurationSeconds{
			value: 10,
		},
	}
}

func (c VideoDescriptionsGopDurationSeconds) Value() int32 {
	return c.value
}

func (c VideoDescriptionsGopDurationSeconds) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoDescriptionsGopDurationSeconds) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
