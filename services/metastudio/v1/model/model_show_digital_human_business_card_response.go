package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDigitalHumanBusinessCardResponse Response Object
type ShowDigitalHumanBusinessCardResponse struct {
	JobInfo *DigitalHumanBusinessCardJobInfo `json:"job_info,omitempty"`

	// 数字人名片模板资产ID。
	CardTempletAssetId *string `json:"card_templet_asset_id,omitempty"`

	CardTextConfig *BusinessCardTextConfig `json:"card_text_config,omitempty"`

	CardImageUrl *BusinessCardImageUrl `json:"card_image_url,omitempty"`

	// 自我介绍文本，用于驱动数字人口型。
	IntroductionText *string `json:"introduction_text,omitempty"`

	// 音色资产ID。
	VoiceAssetId *string `json:"voice_asset_id,omitempty"`

	// 性别。 * MALE：男性 * FEMALE：女性
	Gender *ShowDigitalHumanBusinessCardResponseGender `json:"gender,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDigitalHumanBusinessCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDigitalHumanBusinessCardResponse struct{}"
	}

	return strings.Join([]string{"ShowDigitalHumanBusinessCardResponse", string(data)}, " ")
}

type ShowDigitalHumanBusinessCardResponseGender struct {
	value string
}

type ShowDigitalHumanBusinessCardResponseGenderEnum struct {
	MALE   ShowDigitalHumanBusinessCardResponseGender
	FEMALE ShowDigitalHumanBusinessCardResponseGender
}

func GetShowDigitalHumanBusinessCardResponseGenderEnum() ShowDigitalHumanBusinessCardResponseGenderEnum {
	return ShowDigitalHumanBusinessCardResponseGenderEnum{
		MALE: ShowDigitalHumanBusinessCardResponseGender{
			value: "MALE",
		},
		FEMALE: ShowDigitalHumanBusinessCardResponseGender{
			value: "FEMALE",
		},
	}
}

func (c ShowDigitalHumanBusinessCardResponseGender) Value() string {
	return c.value
}

func (c ShowDigitalHumanBusinessCardResponseGender) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDigitalHumanBusinessCardResponseGender) UnmarshalJSON(b []byte) error {
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
