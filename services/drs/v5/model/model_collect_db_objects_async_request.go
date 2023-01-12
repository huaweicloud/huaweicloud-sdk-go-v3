package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CollectDbObjectsAsyncRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *CollectDbObjectsAsyncRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`

	// 查询对象信息类型。取值： - source：查询源库对象信息。 - modified：查询已选择的（已同步的和未下发的）对象信息。 - synchronized：查询已同步的（已下发的）对象信息 ， 使用场景在任务处于全量中或者增量中。
	Type CollectDbObjectsAsyncRequestType `json:"type"`

	// 是否强制刷新。取值： - true：是，表示从源库重新查询。    - false：否，表示从已缓存中数据查询。
	IsRefresh *bool `json:"is_refresh,omitempty"`

	// 查询指定库的信息。
	DbNames *[]string `json:"db_names,omitempty"`
}

func (o CollectDbObjectsAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectDbObjectsAsyncRequest struct{}"
	}

	return strings.Join([]string{"CollectDbObjectsAsyncRequest", string(data)}, " ")
}

type CollectDbObjectsAsyncRequestXLanguage struct {
	value string
}

type CollectDbObjectsAsyncRequestXLanguageEnum struct {
	EN_US CollectDbObjectsAsyncRequestXLanguage
	ZH_CN CollectDbObjectsAsyncRequestXLanguage
}

func GetCollectDbObjectsAsyncRequestXLanguageEnum() CollectDbObjectsAsyncRequestXLanguageEnum {
	return CollectDbObjectsAsyncRequestXLanguageEnum{
		EN_US: CollectDbObjectsAsyncRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CollectDbObjectsAsyncRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CollectDbObjectsAsyncRequestXLanguage) Value() string {
	return c.value
}

func (c CollectDbObjectsAsyncRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectDbObjectsAsyncRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type CollectDbObjectsAsyncRequestType struct {
	value string
}

type CollectDbObjectsAsyncRequestTypeEnum struct {
	SOURCE       CollectDbObjectsAsyncRequestType
	MODIFIED     CollectDbObjectsAsyncRequestType
	SYNCHRONIZED CollectDbObjectsAsyncRequestType
}

func GetCollectDbObjectsAsyncRequestTypeEnum() CollectDbObjectsAsyncRequestTypeEnum {
	return CollectDbObjectsAsyncRequestTypeEnum{
		SOURCE: CollectDbObjectsAsyncRequestType{
			value: "source",
		},
		MODIFIED: CollectDbObjectsAsyncRequestType{
			value: "modified",
		},
		SYNCHRONIZED: CollectDbObjectsAsyncRequestType{
			value: "synchronized",
		},
	}
}

func (c CollectDbObjectsAsyncRequestType) Value() string {
	return c.value
}

func (c CollectDbObjectsAsyncRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectDbObjectsAsyncRequestType) UnmarshalJSON(b []byte) error {
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
