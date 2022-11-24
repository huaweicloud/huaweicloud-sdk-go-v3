package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ThirdActionTemplateItem struct {

	// 名称。
	Name *string `json:"name,omitempty"`

	// 算子分类。默认分类为FileProcess,MediaProcess,ImageProcess,ContentReview,NotificationProcess,VoiceInteraction
	Category *ThirdActionTemplateItemCategory `json:"category,omitempty"`

	RegisterStatus *PublicTemplateRegisterType `json:"register_status,omitempty"`

	// 算子提供方
	ProviderName *string `json:"provider_name,omitempty"`

	// 算子描述
	Description *string `json:"description,omitempty"`

	// 算子创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后修改时间
	LastModifyTime *string `json:"last_modify_time,omitempty"`
}

func (o ThirdActionTemplateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdActionTemplateItem struct{}"
	}

	return strings.Join([]string{"ThirdActionTemplateItem", string(data)}, " ")
}

type ThirdActionTemplateItemCategory struct {
	value string
}

type ThirdActionTemplateItemCategoryEnum struct {
	FILE_PROCESS         ThirdActionTemplateItemCategory
	MEDIA_PROCESS        ThirdActionTemplateItemCategory
	IMAGE_PROCESS        ThirdActionTemplateItemCategory
	CONTENT_REVIEW       ThirdActionTemplateItemCategory
	NOTIFICATION_PROCESS ThirdActionTemplateItemCategory
	VOICE_INTERACTION    ThirdActionTemplateItemCategory
}

func GetThirdActionTemplateItemCategoryEnum() ThirdActionTemplateItemCategoryEnum {
	return ThirdActionTemplateItemCategoryEnum{
		FILE_PROCESS: ThirdActionTemplateItemCategory{
			value: "FileProcess",
		},
		MEDIA_PROCESS: ThirdActionTemplateItemCategory{
			value: "MediaProcess",
		},
		IMAGE_PROCESS: ThirdActionTemplateItemCategory{
			value: "ImageProcess",
		},
		CONTENT_REVIEW: ThirdActionTemplateItemCategory{
			value: "ContentReview",
		},
		NOTIFICATION_PROCESS: ThirdActionTemplateItemCategory{
			value: "NotificationProcess",
		},
		VOICE_INTERACTION: ThirdActionTemplateItemCategory{
			value: "VoiceInteraction",
		},
	}
}

func (c ThirdActionTemplateItemCategory) Value() string {
	return c.value
}

func (c ThirdActionTemplateItemCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThirdActionTemplateItemCategory) UnmarshalJSON(b []byte) error {
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
