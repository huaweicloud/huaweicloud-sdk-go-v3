package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceTagsRequest Request Object
type ShowInstanceTagsRequest struct {

	// 资源类型。 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：仿真回放
	ResourceType ShowInstanceTagsRequestResourceType `json:"resource_type"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowInstanceTagsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceTagsRequest", string(data)}, " ")
}

type ShowInstanceTagsRequestResourceType struct {
	value string
}

type ShowInstanceTagsRequestResourceTypeEnum struct {
	MIGRATION        ShowInstanceTagsRequestResourceType
	SYNC             ShowInstanceTagsRequestResourceType
	CLOUD_DATA_GUARD ShowInstanceTagsRequestResourceType
	SUBSCRIPTION     ShowInstanceTagsRequestResourceType
	BACKUP_MIGRATION ShowInstanceTagsRequestResourceType
	REPLAY           ShowInstanceTagsRequestResourceType
}

func GetShowInstanceTagsRequestResourceTypeEnum() ShowInstanceTagsRequestResourceTypeEnum {
	return ShowInstanceTagsRequestResourceTypeEnum{
		MIGRATION: ShowInstanceTagsRequestResourceType{
			value: "migration",
		},
		SYNC: ShowInstanceTagsRequestResourceType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ShowInstanceTagsRequestResourceType{
			value: "cloudDataGuard",
		},
		SUBSCRIPTION: ShowInstanceTagsRequestResourceType{
			value: "subscription",
		},
		BACKUP_MIGRATION: ShowInstanceTagsRequestResourceType{
			value: "backupMigration",
		},
		REPLAY: ShowInstanceTagsRequestResourceType{
			value: "replay",
		},
	}
}

func (c ShowInstanceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ShowInstanceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ShowInstanceTagsRequestXLanguage struct {
	value string
}

type ShowInstanceTagsRequestXLanguageEnum struct {
	EN_US ShowInstanceTagsRequestXLanguage
	ZH_CN ShowInstanceTagsRequestXLanguage
}

func GetShowInstanceTagsRequestXLanguageEnum() ShowInstanceTagsRequestXLanguageEnum {
	return ShowInstanceTagsRequestXLanguageEnum{
		EN_US: ShowInstanceTagsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowInstanceTagsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowInstanceTagsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowInstanceTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
