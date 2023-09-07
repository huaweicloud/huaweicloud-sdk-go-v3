package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectTagsRequest Request Object
type ListProjectTagsRequest struct {

	// 资源类型。 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：仿真回放
	ResourceType ListProjectTagsRequestResourceType `json:"resource_type"`

	// 请求语言类型。
	XLanguage *ListProjectTagsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}

type ListProjectTagsRequestResourceType struct {
	value string
}

type ListProjectTagsRequestResourceTypeEnum struct {
	MIGRATION        ListProjectTagsRequestResourceType
	SYNC             ListProjectTagsRequestResourceType
	CLOUD_DATA_GUARD ListProjectTagsRequestResourceType
	SUBSCRIPTION     ListProjectTagsRequestResourceType
	BACKUP_MIGRATION ListProjectTagsRequestResourceType
	REPLAY           ListProjectTagsRequestResourceType
}

func GetListProjectTagsRequestResourceTypeEnum() ListProjectTagsRequestResourceTypeEnum {
	return ListProjectTagsRequestResourceTypeEnum{
		MIGRATION: ListProjectTagsRequestResourceType{
			value: "migration",
		},
		SYNC: ListProjectTagsRequestResourceType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ListProjectTagsRequestResourceType{
			value: "cloudDataGuard",
		},
		SUBSCRIPTION: ListProjectTagsRequestResourceType{
			value: "subscription",
		},
		BACKUP_MIGRATION: ListProjectTagsRequestResourceType{
			value: "backupMigration",
		},
		REPLAY: ListProjectTagsRequestResourceType{
			value: "replay",
		},
	}
}

func (c ListProjectTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListProjectTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListProjectTagsRequestXLanguage struct {
	value string
}

type ListProjectTagsRequestXLanguageEnum struct {
	EN_US ListProjectTagsRequestXLanguage
	ZH_CN ListProjectTagsRequestXLanguage
}

func GetListProjectTagsRequestXLanguageEnum() ListProjectTagsRequestXLanguageEnum {
	return ListProjectTagsRequestXLanguageEnum{
		EN_US: ListProjectTagsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListProjectTagsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListProjectTagsRequestXLanguage) Value() string {
	return c.value
}

func (c ListProjectTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
