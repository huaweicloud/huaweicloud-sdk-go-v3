package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowObjectColumnInfoRequest Request Object
type ShowObjectColumnInfoRequest struct {

	// 请求语言类型。
	XLanguage *ShowObjectColumnInfoRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对象id。不填时默认查询库级的对象。
	ObjectId *string `json:"object_id,omitempty"`

	// 是否强制刷新查询结果
	IsRefresh *bool `json:"is_refresh,omitempty"`

	// 查询返回记录的数量限制，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录，默认为0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowObjectColumnInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObjectColumnInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowObjectColumnInfoRequest", string(data)}, " ")
}

type ShowObjectColumnInfoRequestXLanguage struct {
	value string
}

type ShowObjectColumnInfoRequestXLanguageEnum struct {
	EN_US ShowObjectColumnInfoRequestXLanguage
	ZH_CN ShowObjectColumnInfoRequestXLanguage
}

func GetShowObjectColumnInfoRequestXLanguageEnum() ShowObjectColumnInfoRequestXLanguageEnum {
	return ShowObjectColumnInfoRequestXLanguageEnum{
		EN_US: ShowObjectColumnInfoRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowObjectColumnInfoRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowObjectColumnInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ShowObjectColumnInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowObjectColumnInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
