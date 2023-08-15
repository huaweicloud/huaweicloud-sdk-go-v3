package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PublicTemplateItem struct {

	// 名称。
	Name *string `json:"name,omitempty"`

	// 分类。默认分类为FileProcess,MediaProcess,ImageProcess,ContentReview,NotificationProcess,VoiceInteraction
	Category *PublicTemplateItemCategory `json:"category,omitempty"`

	RegisterStatus *PublicTemplateRegisterType `json:"register_status,omitempty"`

	// 提供方
	ProviderName *string `json:"provider_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后修改时间
	LastModifyTime *string `json:"last_modify_time,omitempty"`
}

func (o PublicTemplateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicTemplateItem struct{}"
	}

	return strings.Join([]string{"PublicTemplateItem", string(data)}, " ")
}

type PublicTemplateItemCategory struct {
	value string
}

type PublicTemplateItemCategoryEnum struct {
	FILE_PROCESS         PublicTemplateItemCategory
	MEDIA_PROCESS        PublicTemplateItemCategory
	IMAGE_PROCESS        PublicTemplateItemCategory
	CONTENT_REVIEW       PublicTemplateItemCategory
	NOTIFICATION_PROCESS PublicTemplateItemCategory
	VOICE_INTERACTION    PublicTemplateItemCategory
}

func GetPublicTemplateItemCategoryEnum() PublicTemplateItemCategoryEnum {
	return PublicTemplateItemCategoryEnum{
		FILE_PROCESS: PublicTemplateItemCategory{
			value: "FileProcess",
		},
		MEDIA_PROCESS: PublicTemplateItemCategory{
			value: "MediaProcess",
		},
		IMAGE_PROCESS: PublicTemplateItemCategory{
			value: "ImageProcess",
		},
		CONTENT_REVIEW: PublicTemplateItemCategory{
			value: "ContentReview",
		},
		NOTIFICATION_PROCESS: PublicTemplateItemCategory{
			value: "NotificationProcess",
		},
		VOICE_INTERACTION: PublicTemplateItemCategory{
			value: "VoiceInteraction",
		},
	}
}

func (c PublicTemplateItemCategory) Value() string {
	return c.value
}

func (c PublicTemplateItemCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicTemplateItemCategory) UnmarshalJSON(b []byte) error {
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
