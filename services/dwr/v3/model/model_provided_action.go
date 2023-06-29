package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProvidedAction struct {

	// 算子模板名称。
	TemplateName string `json:"template_name"`

	// 分类。默认分类为FileProcess,MediaProcess,ImageProcess,ContentReview,NotificationProcess,VoiceInteraction
	Category ProvidedActionCategory `json:"category"`

	// 创建时间。
	CreateTime string `json:"create_time"`

	// 最近修改时间。
	LastModifyTime *string `json:"last_modify_time,omitempty"`

	// 可修改参数定义列表。
	Inputs *[]Input `json:"inputs,omitempty"`

	// 可修改参数引用
	DynamicSourceDefinition map[string]interface{} `json:"dynamic_source_definition"`

	// 需要的权限。
	NeedPolicy []Policy `json:"need_policy"`

	// 算子提供方名称。
	ProviderName string `json:"provider_name"`

	// 是否上传了算子包
	IsUploadedFuncPkg bool `json:"is_uploaded_func_pkg"`

	// 上传算子包的临时签名URL地址，用于上传算子包。
	FuncPkgEndpoint string `json:"func_pkg_endpoint"`

	// 上传算子包的大小。小于100M
	UploadFuncPkgSize *int32 `json:"upload_func_pkg_size,omitempty"`

	// 上传算子包的etag。
	UploadFuncPkgEtag *string `json:"upload_func_pkg_etag,omitempty"`

	RegisterStatus *PublicTemplateRegisterType `json:"register_status"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 函数URN
	FunctionTemplate string `json:"function_template"`

	// 算子提供方IAM DomainID
	ProviderDomainid string `json:"provider_domainid"`

	// 算子提供方IAM UserID
	ProviderUserid *string `json:"provider_userid,omitempty"`
}

func (o ProvidedAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProvidedAction struct{}"
	}

	return strings.Join([]string{"ProvidedAction", string(data)}, " ")
}

type ProvidedActionCategory struct {
	value string
}

type ProvidedActionCategoryEnum struct {
	FILE_PROCESS         ProvidedActionCategory
	MEDIA_PROCESS        ProvidedActionCategory
	IMAGE_PROCESS        ProvidedActionCategory
	CONTENT_REVIEW       ProvidedActionCategory
	NOTIFICATION_PROCESS ProvidedActionCategory
	VOICE_INTERACTION    ProvidedActionCategory
}

func GetProvidedActionCategoryEnum() ProvidedActionCategoryEnum {
	return ProvidedActionCategoryEnum{
		FILE_PROCESS: ProvidedActionCategory{
			value: "FileProcess",
		},
		MEDIA_PROCESS: ProvidedActionCategory{
			value: "MediaProcess",
		},
		IMAGE_PROCESS: ProvidedActionCategory{
			value: "ImageProcess",
		},
		CONTENT_REVIEW: ProvidedActionCategory{
			value: "ContentReview",
		},
		NOTIFICATION_PROCESS: ProvidedActionCategory{
			value: "NotificationProcess",
		},
		VOICE_INTERACTION: ProvidedActionCategory{
			value: "VoiceInteraction",
		},
	}
}

func (c ProvidedActionCategory) Value() string {
	return c.value
}

func (c ProvidedActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProvidedActionCategory) UnmarshalJSON(b []byte) error {
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
