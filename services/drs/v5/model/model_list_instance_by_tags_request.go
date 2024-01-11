package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceByTagsRequest Request Object
type ListInstanceByTagsRequest struct {

	// 资源类型。 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：录制回放
	ResourceType ListInstanceByTagsRequestResourceType `json:"resource_type"`

	// 请求语言类型。
	XLanguage *ListInstanceByTagsRequestXLanguage `json:"X-Language,omitempty"`

	// 查询记录数，默认为1000，limit最多为1000,不能为负数，最小值为1。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	Body *QueryInstanceByTagReq `json:"body,omitempty"`
}

func (o ListInstanceByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceByTagsRequest", string(data)}, " ")
}

type ListInstanceByTagsRequestResourceType struct {
	value string
}

type ListInstanceByTagsRequestResourceTypeEnum struct {
	MIGRATION        ListInstanceByTagsRequestResourceType
	SYNC             ListInstanceByTagsRequestResourceType
	CLOUD_DATA_GUARD ListInstanceByTagsRequestResourceType
	SUBSCRIPTION     ListInstanceByTagsRequestResourceType
	BACKUP_MIGRATION ListInstanceByTagsRequestResourceType
	REPLAY           ListInstanceByTagsRequestResourceType
}

func GetListInstanceByTagsRequestResourceTypeEnum() ListInstanceByTagsRequestResourceTypeEnum {
	return ListInstanceByTagsRequestResourceTypeEnum{
		MIGRATION: ListInstanceByTagsRequestResourceType{
			value: "migration",
		},
		SYNC: ListInstanceByTagsRequestResourceType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ListInstanceByTagsRequestResourceType{
			value: "cloudDataGuard",
		},
		SUBSCRIPTION: ListInstanceByTagsRequestResourceType{
			value: "subscription",
		},
		BACKUP_MIGRATION: ListInstanceByTagsRequestResourceType{
			value: "backupMigration",
		},
		REPLAY: ListInstanceByTagsRequestResourceType{
			value: "replay",
		},
	}
}

func (c ListInstanceByTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListInstanceByTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceByTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListInstanceByTagsRequestXLanguage struct {
	value string
}

type ListInstanceByTagsRequestXLanguageEnum struct {
	EN_US ListInstanceByTagsRequestXLanguage
	ZH_CN ListInstanceByTagsRequestXLanguage
}

func GetListInstanceByTagsRequestXLanguageEnum() ListInstanceByTagsRequestXLanguageEnum {
	return ListInstanceByTagsRequestXLanguageEnum{
		EN_US: ListInstanceByTagsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListInstanceByTagsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListInstanceByTagsRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceByTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceByTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
