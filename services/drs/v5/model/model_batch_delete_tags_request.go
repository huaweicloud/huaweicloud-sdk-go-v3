package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteTagsRequest Request Object
type BatchDeleteTagsRequest struct {

	// 资源类型。 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：录制回放
	ResourceType BatchDeleteTagsRequestResourceType `json:"resource_type"`

	// 请求语言类型。
	XLanguage *BatchDeleteTagsRequestXLanguage `json:"X-Language,omitempty"`

	// 资源ID，即DRS任务ID。
	ResourceId string `json:"resource_id"`

	Body *BatchDeleteTagReq `json:"body,omitempty"`
}

func (o BatchDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequest", string(data)}, " ")
}

type BatchDeleteTagsRequestResourceType struct {
	value string
}

type BatchDeleteTagsRequestResourceTypeEnum struct {
	MIGRATION        BatchDeleteTagsRequestResourceType
	SYNC             BatchDeleteTagsRequestResourceType
	CLOUD_DATA_GUARD BatchDeleteTagsRequestResourceType
	SUBSCRIPTION     BatchDeleteTagsRequestResourceType
	BACKUP_MIGRATION BatchDeleteTagsRequestResourceType
	REPLAY           BatchDeleteTagsRequestResourceType
}

func GetBatchDeleteTagsRequestResourceTypeEnum() BatchDeleteTagsRequestResourceTypeEnum {
	return BatchDeleteTagsRequestResourceTypeEnum{
		MIGRATION: BatchDeleteTagsRequestResourceType{
			value: "migration",
		},
		SYNC: BatchDeleteTagsRequestResourceType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: BatchDeleteTagsRequestResourceType{
			value: "cloudDataGuard",
		},
		SUBSCRIPTION: BatchDeleteTagsRequestResourceType{
			value: "subscription",
		},
		BACKUP_MIGRATION: BatchDeleteTagsRequestResourceType{
			value: "backupMigration",
		},
		REPLAY: BatchDeleteTagsRequestResourceType{
			value: "replay",
		},
	}
}

func (c BatchDeleteTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchDeleteTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type BatchDeleteTagsRequestXLanguage struct {
	value string
}

type BatchDeleteTagsRequestXLanguageEnum struct {
	EN_US BatchDeleteTagsRequestXLanguage
	ZH_CN BatchDeleteTagsRequestXLanguage
}

func GetBatchDeleteTagsRequestXLanguageEnum() BatchDeleteTagsRequestXLanguageEnum {
	return BatchDeleteTagsRequestXLanguageEnum{
		EN_US: BatchDeleteTagsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchDeleteTagsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchDeleteTagsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchDeleteTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
