package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDigitalHumanBusinessCardReq 数字人名片制作创建请求。
type CreateDigitalHumanBusinessCardReq struct {

	// 数字人名片类型。 * 2D_DIGITAL_HUMAN_CARD：分身数字人名片。
	BusinessCardType CreateDigitalHumanBusinessCardReqBusinessCardType `json:"business_card_type"`

	// 数字人名片模板资产ID。
	CardTempletAssetId string `json:"card_templet_asset_id"`

	CardTextConfig *BusinessCardTextConfig `json:"card_text_config"`

	CardImageConfig *BusinessCardImageConfig `json:"card_image_config"`

	// 自我介绍文本，用于驱动数字人口型。
	IntroductionText string `json:"introduction_text"`

	// 音色资产ID。
	VoiceAssetId string `json:"voice_asset_id"`

	// 输出名片视频资产名称。默认取card_name的值
	VideoAssetName *string `json:"video_asset_name,omitempty"`

	// 性别。 * MALE：男性 * FEMALE：女性
	Gender *CreateDigitalHumanBusinessCardReqGender `json:"gender,omitempty"`
}

func (o CreateDigitalHumanBusinessCardReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDigitalHumanBusinessCardReq struct{}"
	}

	return strings.Join([]string{"CreateDigitalHumanBusinessCardReq", string(data)}, " ")
}

type CreateDigitalHumanBusinessCardReqBusinessCardType struct {
	value string
}

type CreateDigitalHumanBusinessCardReqBusinessCardTypeEnum struct {
	E_2_D_DIGITAL_HUMAN_CARD CreateDigitalHumanBusinessCardReqBusinessCardType
}

func GetCreateDigitalHumanBusinessCardReqBusinessCardTypeEnum() CreateDigitalHumanBusinessCardReqBusinessCardTypeEnum {
	return CreateDigitalHumanBusinessCardReqBusinessCardTypeEnum{
		E_2_D_DIGITAL_HUMAN_CARD: CreateDigitalHumanBusinessCardReqBusinessCardType{
			value: "2D_DIGITAL_HUMAN_CARD",
		},
	}
}

func (c CreateDigitalHumanBusinessCardReqBusinessCardType) Value() string {
	return c.value
}

func (c CreateDigitalHumanBusinessCardReqBusinessCardType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDigitalHumanBusinessCardReqBusinessCardType) UnmarshalJSON(b []byte) error {
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

type CreateDigitalHumanBusinessCardReqGender struct {
	value string
}

type CreateDigitalHumanBusinessCardReqGenderEnum struct {
	MALE   CreateDigitalHumanBusinessCardReqGender
	FEMALE CreateDigitalHumanBusinessCardReqGender
}

func GetCreateDigitalHumanBusinessCardReqGenderEnum() CreateDigitalHumanBusinessCardReqGenderEnum {
	return CreateDigitalHumanBusinessCardReqGenderEnum{
		MALE: CreateDigitalHumanBusinessCardReqGender{
			value: "MALE",
		},
		FEMALE: CreateDigitalHumanBusinessCardReqGender{
			value: "FEMALE",
		},
	}
}

func (c CreateDigitalHumanBusinessCardReqGender) Value() string {
	return c.value
}

func (c CreateDigitalHumanBusinessCardReqGender) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDigitalHumanBusinessCardReqGender) UnmarshalJSON(b []byte) error {
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
