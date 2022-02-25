package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SearchQosParticipantDetailRequest struct {
	// 会议UUID。最大不超过64个字节。

	ConfUUID string `json:"confUUID"`
	// 会议类别。 * online：在线会议，在召开的会议。 * history：历史会议，已召开的会议。

	ConfType SearchQosParticipantDetailRequestConfType `json:"confType"`
	// 与会者ID。最大不超过64个字节。

	ParticipantID string `json:"participantID"`
	// Qos类型。 - audio：音频。 - video：视频。 - screen：屏幕共享。 - cpu：cpu。

	QosType SearchQosParticipantDetailRequestQosType `json:"qosType"`
}

func (o SearchQosParticipantDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosParticipantDetailRequest struct{}"
	}

	return strings.Join([]string{"SearchQosParticipantDetailRequest", string(data)}, " ")
}

type SearchQosParticipantDetailRequestConfType struct {
	value string
}

type SearchQosParticipantDetailRequestConfTypeEnum struct {
	ONLINE  SearchQosParticipantDetailRequestConfType
	HISTORY SearchQosParticipantDetailRequestConfType
}

func GetSearchQosParticipantDetailRequestConfTypeEnum() SearchQosParticipantDetailRequestConfTypeEnum {
	return SearchQosParticipantDetailRequestConfTypeEnum{
		ONLINE: SearchQosParticipantDetailRequestConfType{
			value: "online",
		},
		HISTORY: SearchQosParticipantDetailRequestConfType{
			value: "history",
		},
	}
}

func (c SearchQosParticipantDetailRequestConfType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchQosParticipantDetailRequestConfType) UnmarshalJSON(b []byte) error {
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

type SearchQosParticipantDetailRequestQosType struct {
	value string
}

type SearchQosParticipantDetailRequestQosTypeEnum struct {
	AUDIO  SearchQosParticipantDetailRequestQosType
	VIDEO  SearchQosParticipantDetailRequestQosType
	SCREEN SearchQosParticipantDetailRequestQosType
	CPU    SearchQosParticipantDetailRequestQosType
}

func GetSearchQosParticipantDetailRequestQosTypeEnum() SearchQosParticipantDetailRequestQosTypeEnum {
	return SearchQosParticipantDetailRequestQosTypeEnum{
		AUDIO: SearchQosParticipantDetailRequestQosType{
			value: "audio",
		},
		VIDEO: SearchQosParticipantDetailRequestQosType{
			value: "video",
		},
		SCREEN: SearchQosParticipantDetailRequestQosType{
			value: "screen",
		},
		CPU: SearchQosParticipantDetailRequestQosType{
			value: "cpu",
		},
	}
}

func (c SearchQosParticipantDetailRequestQosType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchQosParticipantDetailRequestQosType) UnmarshalJSON(b []byte) error {
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
