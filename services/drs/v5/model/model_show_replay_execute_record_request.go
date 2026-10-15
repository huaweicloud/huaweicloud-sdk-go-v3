package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReplayExecuteRecordRequest Request Object
type ShowReplayExecuteRecordRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowReplayExecuteRecordRequestXLanguage `json:"X-Language,omitempty"`

	// SQL执行记录的起始时间，格式为秒级时间戳。
	StartTime int64 `json:"start_time"`

	// SQL执行记录的结束时间，格式为秒级时间戳。
	EndTime int64 `json:"end_time"`

	// 偏移量，表示查询该偏移量后面的记录，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制，默认为10，最大值为2000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowReplayExecuteRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayExecuteRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowReplayExecuteRecordRequest", string(data)}, " ")
}

type ShowReplayExecuteRecordRequestXLanguage struct {
	value string
}

type ShowReplayExecuteRecordRequestXLanguageEnum struct {
	EN_US ShowReplayExecuteRecordRequestXLanguage
	ZH_CN ShowReplayExecuteRecordRequestXLanguage
}

func GetShowReplayExecuteRecordRequestXLanguageEnum() ShowReplayExecuteRecordRequestXLanguageEnum {
	return ShowReplayExecuteRecordRequestXLanguageEnum{
		EN_US: ShowReplayExecuteRecordRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowReplayExecuteRecordRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowReplayExecuteRecordRequestXLanguage) Value() string {
	return c.value
}

func (c ShowReplayExecuteRecordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplayExecuteRecordRequestXLanguage) UnmarshalJSON(b []byte) error {
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
