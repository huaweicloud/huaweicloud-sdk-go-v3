package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowComparePolicyResponse Response Object
type ShowComparePolicyResponse struct {

	// 对比时间。
	Period *string `json:"period,omitempty"`

	// 对比策略状态。 - OPEN：开启 - CLOSED：关闭 - NO_SUPPORT：不支持
	Status *ShowComparePolicyResponseStatus `json:"status,omitempty"`

	// 对比开始时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// 对比结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 对比类型： - object_comparison：对象对比。 - lines：行对比。 - account：用户对比。
	CompareType *[]string `json:"compare_type,omitempty"`

	// 下次对比时间，UTC时间，例如：2023-06-12T08:00:00Z
	NextCompareTime *string `json:"next_compare_time,omitempty"`

	// 对比策略。 - normal：普通对比 - manyToOne：多对一对比
	ComparePolicy *string `json:"compare_policy,omitempty"`

	// 间隔时间。
	IntervalHour   *int32 `json:"interval_hour,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowComparePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComparePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowComparePolicyResponse", string(data)}, " ")
}

type ShowComparePolicyResponseStatus struct {
	value string
}

type ShowComparePolicyResponseStatusEnum struct {
	OPEN       ShowComparePolicyResponseStatus
	CLOSED     ShowComparePolicyResponseStatus
	NO_SUPPORT ShowComparePolicyResponseStatus
}

func GetShowComparePolicyResponseStatusEnum() ShowComparePolicyResponseStatusEnum {
	return ShowComparePolicyResponseStatusEnum{
		OPEN: ShowComparePolicyResponseStatus{
			value: "OPEN",
		},
		CLOSED: ShowComparePolicyResponseStatus{
			value: "CLOSED",
		},
		NO_SUPPORT: ShowComparePolicyResponseStatus{
			value: "NO_SUPPORT",
		},
	}
}

func (c ShowComparePolicyResponseStatus) Value() string {
	return c.value
}

func (c ShowComparePolicyResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowComparePolicyResponseStatus) UnmarshalJSON(b []byte) error {
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
