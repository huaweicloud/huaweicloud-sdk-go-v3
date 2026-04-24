package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCompareUsersDetailRequest Request Object
type ShowCompareUsersDetailRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务的ID。
	CompareJobId string `json:"compare_job_id"`

	// 请求语言类型。
	XLanguage *ShowCompareUsersDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录，默认为0。
	Offset *string `json:"offset,omitempty"`

	// 查询返回记录的数量限制，默认为10。
	Limit *string `json:"limit,omitempty"`
}

func (o ShowCompareUsersDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompareUsersDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCompareUsersDetailRequest", string(data)}, " ")
}

type ShowCompareUsersDetailRequestXLanguage struct {
	value string
}

type ShowCompareUsersDetailRequestXLanguageEnum struct {
	EN_US ShowCompareUsersDetailRequestXLanguage
	ZH_CN ShowCompareUsersDetailRequestXLanguage
}

func GetShowCompareUsersDetailRequestXLanguageEnum() ShowCompareUsersDetailRequestXLanguageEnum {
	return ShowCompareUsersDetailRequestXLanguageEnum{
		EN_US: ShowCompareUsersDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowCompareUsersDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowCompareUsersDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowCompareUsersDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCompareUsersDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
