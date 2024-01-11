package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateTagsRequest Request Object
type BatchCreateTagsRequest struct {

	// 资源类型。 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：录制回放
	ResourceType BatchCreateTagsRequestResourceType `json:"resource_type"`

	// 请求语言类型。
	XLanguage *BatchCreateTagsRequestXLanguage `json:"X-Language,omitempty"`

	// 资源ID，即DRS任务ID。
	ResourceId string `json:"resource_id"`

	Body *BatchAddTagReq `json:"body,omitempty"`
}

func (o BatchCreateTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateTagsRequest", string(data)}, " ")
}

type BatchCreateTagsRequestResourceType struct {
	value string
}

type BatchCreateTagsRequestResourceTypeEnum struct {
	MIGRATION        BatchCreateTagsRequestResourceType
	SYNC             BatchCreateTagsRequestResourceType
	CLOUD_DATA_GUARD BatchCreateTagsRequestResourceType
	SUBSCRIPTION     BatchCreateTagsRequestResourceType
	BACKUP_MIGRATION BatchCreateTagsRequestResourceType
	REPLAY           BatchCreateTagsRequestResourceType
}

func GetBatchCreateTagsRequestResourceTypeEnum() BatchCreateTagsRequestResourceTypeEnum {
	return BatchCreateTagsRequestResourceTypeEnum{
		MIGRATION: BatchCreateTagsRequestResourceType{
			value: "migration",
		},
		SYNC: BatchCreateTagsRequestResourceType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: BatchCreateTagsRequestResourceType{
			value: "cloudDataGuard",
		},
		SUBSCRIPTION: BatchCreateTagsRequestResourceType{
			value: "subscription",
		},
		BACKUP_MIGRATION: BatchCreateTagsRequestResourceType{
			value: "backupMigration",
		},
		REPLAY: BatchCreateTagsRequestResourceType{
			value: "replay",
		},
	}
}

func (c BatchCreateTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchCreateTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type BatchCreateTagsRequestXLanguage struct {
	value string
}

type BatchCreateTagsRequestXLanguageEnum struct {
	EN_US BatchCreateTagsRequestXLanguage
	ZH_CN BatchCreateTagsRequestXLanguage
}

func GetBatchCreateTagsRequestXLanguageEnum() BatchCreateTagsRequestXLanguageEnum {
	return BatchCreateTagsRequestXLanguageEnum{
		EN_US: BatchCreateTagsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchCreateTagsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchCreateTagsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchCreateTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
