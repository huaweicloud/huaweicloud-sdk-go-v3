package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ActionTemplateItem struct {

	// 名称。
	Name *string `json:"name,omitempty"`

	// 第三方算子模板的分类。默认分类为FileProcess,MediaProcess,ImageProcess,ContentReview,NotificationProcess,VoiceInteraction
	Category *ActionTemplateItemCategory `json:"category,omitempty"`
}

func (o ActionTemplateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionTemplateItem struct{}"
	}

	return strings.Join([]string{"ActionTemplateItem", string(data)}, " ")
}

type ActionTemplateItemCategory struct {
	value string
}

type ActionTemplateItemCategoryEnum struct {
	FILE_PROCESS         ActionTemplateItemCategory
	MEDIA_PROCESS        ActionTemplateItemCategory
	IMAGE_PROCESS        ActionTemplateItemCategory
	CONTENT_REVIEW       ActionTemplateItemCategory
	NOTIFICATION_PROCESS ActionTemplateItemCategory
	VOICE_INTERACTION    ActionTemplateItemCategory
}

func GetActionTemplateItemCategoryEnum() ActionTemplateItemCategoryEnum {
	return ActionTemplateItemCategoryEnum{
		FILE_PROCESS: ActionTemplateItemCategory{
			value: "FileProcess",
		},
		MEDIA_PROCESS: ActionTemplateItemCategory{
			value: "MediaProcess",
		},
		IMAGE_PROCESS: ActionTemplateItemCategory{
			value: "ImageProcess",
		},
		CONTENT_REVIEW: ActionTemplateItemCategory{
			value: "ContentReview",
		},
		NOTIFICATION_PROCESS: ActionTemplateItemCategory{
			value: "NotificationProcess",
		},
		VOICE_INTERACTION: ActionTemplateItemCategory{
			value: "VoiceInteraction",
		},
	}
}

func (c ActionTemplateItemCategory) Value() string {
	return c.value
}

func (c ActionTemplateItemCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionTemplateItemCategory) UnmarshalJSON(b []byte) error {
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
