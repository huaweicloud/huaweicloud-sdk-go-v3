package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountInstanceByTagsRequest Request Object
type CountInstanceByTagsRequest struct {

	// 资源类型。 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：录制回放
	ResourceType CountInstanceByTagsRequestResourceType `json:"resource_type"`

	// 请求语言类型。
	XLanguage *CountInstanceByTagsRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryInstanceByTagReq `json:"body,omitempty"`
}

func (o CountInstanceByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInstanceByTagsRequest struct{}"
	}

	return strings.Join([]string{"CountInstanceByTagsRequest", string(data)}, " ")
}

type CountInstanceByTagsRequestResourceType struct {
	value string
}

type CountInstanceByTagsRequestResourceTypeEnum struct {
	MIGRATION        CountInstanceByTagsRequestResourceType
	SYNC             CountInstanceByTagsRequestResourceType
	CLOUD_DATA_GUARD CountInstanceByTagsRequestResourceType
	SUBSCRIPTION     CountInstanceByTagsRequestResourceType
	BACKUP_MIGRATION CountInstanceByTagsRequestResourceType
	REPLAY           CountInstanceByTagsRequestResourceType
}

func GetCountInstanceByTagsRequestResourceTypeEnum() CountInstanceByTagsRequestResourceTypeEnum {
	return CountInstanceByTagsRequestResourceTypeEnum{
		MIGRATION: CountInstanceByTagsRequestResourceType{
			value: "migration",
		},
		SYNC: CountInstanceByTagsRequestResourceType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: CountInstanceByTagsRequestResourceType{
			value: "cloudDataGuard",
		},
		SUBSCRIPTION: CountInstanceByTagsRequestResourceType{
			value: "subscription",
		},
		BACKUP_MIGRATION: CountInstanceByTagsRequestResourceType{
			value: "backupMigration",
		},
		REPLAY: CountInstanceByTagsRequestResourceType{
			value: "replay",
		},
	}
}

func (c CountInstanceByTagsRequestResourceType) Value() string {
	return c.value
}

func (c CountInstanceByTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountInstanceByTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type CountInstanceByTagsRequestXLanguage struct {
	value string
}

type CountInstanceByTagsRequestXLanguageEnum struct {
	EN_US CountInstanceByTagsRequestXLanguage
	ZH_CN CountInstanceByTagsRequestXLanguage
}

func GetCountInstanceByTagsRequestXLanguageEnum() CountInstanceByTagsRequestXLanguageEnum {
	return CountInstanceByTagsRequestXLanguageEnum{
		EN_US: CountInstanceByTagsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CountInstanceByTagsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CountInstanceByTagsRequestXLanguage) Value() string {
	return c.value
}

func (c CountInstanceByTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountInstanceByTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
