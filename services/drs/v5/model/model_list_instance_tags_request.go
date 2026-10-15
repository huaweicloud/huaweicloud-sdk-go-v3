package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceTagsRequest Request Object
type ListInstanceTagsRequest struct {

	// **参数解释**：  DRS资源类型，用于指定需要操作标签的资源所属的业务类型。不同资源类型对应不同的DRS任务类型，查询标签时需根据实际任务类型传入对应值。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：录制回放 - verify：校验任务  **默认取值**：  不涉及
	ResourceType ListInstanceTagsRequestResourceType `json:"resource_type"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文 - zh-cn：中文  **默认取值**：  en-us
	XLanguage *ListInstanceTagsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  资源ID，即DRS任务ID。用于唯一标识一个DRS任务资源，在标签管理接口中用于定位需要查询标签的具体任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ResourceId string `json:"resource_id"`
}

func (o ListInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceTagsRequest", string(data)}, " ")
}

type ListInstanceTagsRequestResourceType struct {
	value string
}

type ListInstanceTagsRequestResourceTypeEnum struct {
	MIGRATION        ListInstanceTagsRequestResourceType
	SYNC             ListInstanceTagsRequestResourceType
	CLOUD_DATA_GUARD ListInstanceTagsRequestResourceType
	SUBSCRIPTION     ListInstanceTagsRequestResourceType
	BACKUP_MIGRATION ListInstanceTagsRequestResourceType
	REPLAY           ListInstanceTagsRequestResourceType
	VERIFY           ListInstanceTagsRequestResourceType
}

func GetListInstanceTagsRequestResourceTypeEnum() ListInstanceTagsRequestResourceTypeEnum {
	return ListInstanceTagsRequestResourceTypeEnum{
		MIGRATION: ListInstanceTagsRequestResourceType{
			value: "migration",
		},
		SYNC: ListInstanceTagsRequestResourceType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ListInstanceTagsRequestResourceType{
			value: "cloudDataGuard",
		},
		SUBSCRIPTION: ListInstanceTagsRequestResourceType{
			value: "subscription",
		},
		BACKUP_MIGRATION: ListInstanceTagsRequestResourceType{
			value: "backupMigration",
		},
		REPLAY: ListInstanceTagsRequestResourceType{
			value: "replay",
		},
		VERIFY: ListInstanceTagsRequestResourceType{
			value: "verify",
		},
	}
}

func (c ListInstanceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListInstanceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListInstanceTagsRequestXLanguage struct {
	value string
}

type ListInstanceTagsRequestXLanguageEnum struct {
	EN_US ListInstanceTagsRequestXLanguage
	ZH_CN ListInstanceTagsRequestXLanguage
}

func GetListInstanceTagsRequestXLanguageEnum() ListInstanceTagsRequestXLanguageEnum {
	return ListInstanceTagsRequestXLanguageEnum{
		EN_US: ListInstanceTagsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListInstanceTagsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListInstanceTagsRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
