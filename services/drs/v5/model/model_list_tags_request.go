package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTagsRequest Request Object
type ListTagsRequest struct {

	// 资源类型。 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：录制回放 - verify：校验任务
	ResourceType ListTagsRequestResourceType `json:"resource_type"`

	// 请求语言类型。
	XLanguage *ListTagsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}

type ListTagsRequestResourceType struct {
	value string
}

type ListTagsRequestResourceTypeEnum struct {
	MIGRATION        ListTagsRequestResourceType
	SYNC             ListTagsRequestResourceType
	CLOUD_DATA_GUARD ListTagsRequestResourceType
	SUBSCRIPTION     ListTagsRequestResourceType
	BACKUP_MIGRATION ListTagsRequestResourceType
	REPLAY           ListTagsRequestResourceType
	VERIFY           ListTagsRequestResourceType
}

func GetListTagsRequestResourceTypeEnum() ListTagsRequestResourceTypeEnum {
	return ListTagsRequestResourceTypeEnum{
		MIGRATION: ListTagsRequestResourceType{
			value: "migration",
		},
		SYNC: ListTagsRequestResourceType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ListTagsRequestResourceType{
			value: "cloudDataGuard",
		},
		SUBSCRIPTION: ListTagsRequestResourceType{
			value: "subscription",
		},
		BACKUP_MIGRATION: ListTagsRequestResourceType{
			value: "backupMigration",
		},
		REPLAY: ListTagsRequestResourceType{
			value: "replay",
		},
		VERIFY: ListTagsRequestResourceType{
			value: "verify",
		},
	}
}

func (c ListTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestXLanguage struct {
	value string
}

type ListTagsRequestXLanguageEnum struct {
	EN_US ListTagsRequestXLanguage
	ZH_CN ListTagsRequestXLanguage
}

func GetListTagsRequestXLanguageEnum() ListTagsRequestXLanguageEnum {
	return ListTagsRequestXLanguageEnum{
		EN_US: ListTagsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListTagsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListTagsRequestXLanguage) Value() string {
	return c.value
}

func (c ListTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
