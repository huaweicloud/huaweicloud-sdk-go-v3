package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IndividualStreamJobReq 单流任务请求，转推和录制至少选一个
type IndividualStreamJobReq struct {

	// 房间id
	RoomId string `json:"room_id"`

	// 选看的用户id，单个录制任务内保证唯一
	UserId string `json:"user_id"`

	//  是否录制音频。  - true：录制音频 - false：不录制音频  缺省为true。
	IsRecordAudio *bool `json:"is_record_audio,omitempty"`

	// 标识视频流的类型，可选摄像头流或者屏幕分享流，未填写表示不录制视频。  - CAMERASTREAM：摄像头视频流 - SCREENSTREAM：屏幕分享视频流  默认为CAMERASTREAM。
	VideoType *IndividualStreamJobReqVideoType `json:"video_type,omitempty"`

	// 指定窗口拉取的分辨率档位。  - LD - SD - HD - FHD  缺省为FHD。
	SelectStreamType *IndividualStreamJobReqSelectStreamType `json:"select_stream_type,omitempty"`

	// 最长空闲频道时间。  取值范围：[5，43200]，默认值为30。  单位：秒。  如果频道内无连麦方的状态持续超过该时间，录制程序会自动退出。退出后，再次调用start请求，会产生新的录制任务。  连麦方指：joiner或者publisher的用户。
	MaxIdleTime *int32 `json:"max_idle_time,omitempty"`

	PublishParam *PublishParam `json:"publish_param,omitempty"`

	RecordParam *RecordParam `json:"record_param,omitempty"`
}

func (o IndividualStreamJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndividualStreamJobReq struct{}"
	}

	return strings.Join([]string{"IndividualStreamJobReq", string(data)}, " ")
}

type IndividualStreamJobReqVideoType struct {
	value string
}

type IndividualStreamJobReqVideoTypeEnum struct {
	CAMERASTREAM IndividualStreamJobReqVideoType
	SCREENSTREAM IndividualStreamJobReqVideoType
}

func GetIndividualStreamJobReqVideoTypeEnum() IndividualStreamJobReqVideoTypeEnum {
	return IndividualStreamJobReqVideoTypeEnum{
		CAMERASTREAM: IndividualStreamJobReqVideoType{
			value: "CAMERASTREAM",
		},
		SCREENSTREAM: IndividualStreamJobReqVideoType{
			value: "SCREENSTREAM",
		},
	}
}

func (c IndividualStreamJobReqVideoType) Value() string {
	return c.value
}

func (c IndividualStreamJobReqVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IndividualStreamJobReqVideoType) UnmarshalJSON(b []byte) error {
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

type IndividualStreamJobReqSelectStreamType struct {
	value string
}

type IndividualStreamJobReqSelectStreamTypeEnum struct {
	LD  IndividualStreamJobReqSelectStreamType
	SD  IndividualStreamJobReqSelectStreamType
	HD  IndividualStreamJobReqSelectStreamType
	FHD IndividualStreamJobReqSelectStreamType
}

func GetIndividualStreamJobReqSelectStreamTypeEnum() IndividualStreamJobReqSelectStreamTypeEnum {
	return IndividualStreamJobReqSelectStreamTypeEnum{
		LD: IndividualStreamJobReqSelectStreamType{
			value: "LD",
		},
		SD: IndividualStreamJobReqSelectStreamType{
			value: "SD",
		},
		HD: IndividualStreamJobReqSelectStreamType{
			value: "HD",
		},
		FHD: IndividualStreamJobReqSelectStreamType{
			value: "FHD",
		},
	}
}

func (c IndividualStreamJobReqSelectStreamType) Value() string {
	return c.value
}

func (c IndividualStreamJobReqSelectStreamType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IndividualStreamJobReqSelectStreamType) UnmarshalJSON(b []byte) error {
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
