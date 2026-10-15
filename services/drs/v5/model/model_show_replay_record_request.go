package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReplayRecordRequest Request Object
type ShowReplayRecordRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowReplayRecordRequestXLanguage `json:"X-Language,omitempty"`

	// SQL执行记录的起始时间，格式为秒级时间戳。
	StartTime int64 `json:"start_time"`

	// SQL执行记录的结束时间，格式为秒级时间戳。
	EndTime int64 `json:"end_time"`

	// 偏移量，表示查询该偏移量后面的记录，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制，默认为10，最大值为2000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowReplayRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowReplayRecordRequest", string(data)}, " ")
}

type ShowReplayRecordRequestXLanguage struct {
	value string
}

type ShowReplayRecordRequestXLanguageEnum struct {
	EN_US ShowReplayRecordRequestXLanguage
	ZH_CN ShowReplayRecordRequestXLanguage
}

func GetShowReplayRecordRequestXLanguageEnum() ShowReplayRecordRequestXLanguageEnum {
	return ShowReplayRecordRequestXLanguageEnum{
		EN_US: ShowReplayRecordRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowReplayRecordRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowReplayRecordRequestXLanguage) Value() string {
	return c.value
}

func (c ShowReplayRecordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplayRecordRequestXLanguage) UnmarshalJSON(b []byte) error {
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
