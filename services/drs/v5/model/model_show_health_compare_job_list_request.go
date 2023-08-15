package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHealthCompareJobListRequest Request Object
type ShowHealthCompareJobListRequest struct {

	// 请求语言类型。
	XLanguage *ShowHealthCompareJobListRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 状态，不传查询所有状态。 - WAITING_FOR_RUNNING：等待启动中 - RUNNING：运行中 - SUCCESSFUL：完成 - FAILED：失败 - CANCELLED：已取消 - TIMEOUT_INTERRUPT：超时中断 - FULL_DOING：全量校验中 - INCRE_DOING：增量校验中
	Status *ShowHealthCompareJobListRequestStatus `json:"status,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。默认为10，取值范围【1-1000】
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowHealthCompareJobListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthCompareJobListRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthCompareJobListRequest", string(data)}, " ")
}

type ShowHealthCompareJobListRequestXLanguage struct {
	value string
}

type ShowHealthCompareJobListRequestXLanguageEnum struct {
	EN_US ShowHealthCompareJobListRequestXLanguage
	ZH_CN ShowHealthCompareJobListRequestXLanguage
}

func GetShowHealthCompareJobListRequestXLanguageEnum() ShowHealthCompareJobListRequestXLanguageEnum {
	return ShowHealthCompareJobListRequestXLanguageEnum{
		EN_US: ShowHealthCompareJobListRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowHealthCompareJobListRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowHealthCompareJobListRequestXLanguage) Value() string {
	return c.value
}

func (c ShowHealthCompareJobListRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHealthCompareJobListRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowHealthCompareJobListRequestStatus struct {
	value string
}

type ShowHealthCompareJobListRequestStatusEnum struct {
	WAITING_FOR_RUNNING ShowHealthCompareJobListRequestStatus
	RUNNING             ShowHealthCompareJobListRequestStatus
	SUCCESSFUL          ShowHealthCompareJobListRequestStatus
	FAILED              ShowHealthCompareJobListRequestStatus
	CANCELLED           ShowHealthCompareJobListRequestStatus
	TIMEOUT_INTERRUPT   ShowHealthCompareJobListRequestStatus
	FULL_DOING          ShowHealthCompareJobListRequestStatus
	INCRE_DOING         ShowHealthCompareJobListRequestStatus
}

func GetShowHealthCompareJobListRequestStatusEnum() ShowHealthCompareJobListRequestStatusEnum {
	return ShowHealthCompareJobListRequestStatusEnum{
		WAITING_FOR_RUNNING: ShowHealthCompareJobListRequestStatus{
			value: "WAITING_FOR_RUNNING",
		},
		RUNNING: ShowHealthCompareJobListRequestStatus{
			value: "RUNNING",
		},
		SUCCESSFUL: ShowHealthCompareJobListRequestStatus{
			value: "SUCCESSFUL",
		},
		FAILED: ShowHealthCompareJobListRequestStatus{
			value: "FAILED",
		},
		CANCELLED: ShowHealthCompareJobListRequestStatus{
			value: "CANCELLED",
		},
		TIMEOUT_INTERRUPT: ShowHealthCompareJobListRequestStatus{
			value: "TIMEOUT_INTERRUPT",
		},
		FULL_DOING: ShowHealthCompareJobListRequestStatus{
			value: "FULL_DOING",
		},
		INCRE_DOING: ShowHealthCompareJobListRequestStatus{
			value: "INCRE_DOING",
		},
	}
}

func (c ShowHealthCompareJobListRequestStatus) Value() string {
	return c.value
}

func (c ShowHealthCompareJobListRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHealthCompareJobListRequestStatus) UnmarshalJSON(b []byte) error {
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
