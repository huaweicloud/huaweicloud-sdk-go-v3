package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyComparePolicyReq 修改对比策略请求体。
type ModifyComparePolicyReq struct {

	// 对比策略开关-open|close。
	Action string `json:"action"`

	// 对比策略周期。 - 每周对比：格式示例：“* * 1,3,5” ，其中“1,3,5”对应星期一、星期三、星期五，根据实际填写。 - 每天对比：固定填写“* * 1,2,3,4,5,6,7”。 - 按小时对比：固定填写“* * 1,2,3,4,5,6,7”。
	Period *string `json:"period,omitempty"`

	// 对比策略生效开始时间，UTC时间，24小时制，时间格式HH:mm:ss，例如：00:00:00，表示UTC时间的00:00:00，北京时间（UTC+08:00）的08:00:00。
	BeginTime *string `json:"begin_time,omitempty"`

	// 对比策略生效结束时间，UTC时间，24小时制，时间格式HH:mm:ss，例如：04:00:00，表示UTC时间的04:00:00，北京时间（UTC+08:00）的12:00:00。
	EndTime *string `json:"end_time,omitempty"`

	// 对比类型列表： - object_comparison：对象对比。 - lines：行对比。 - account：用户对比。
	CompareType *[]string `json:"compare_type,omitempty"`

	// 对比策略： - normal：普通对比。 - manyToOne：多对一对比。
	ComparePolicy *ModifyComparePolicyReqComparePolicy `json:"compare_policy,omitempty"`

	// 间隔时间，按小时对比时填写，表示每隔多久执行一次对比，单位是小时。
	IntervalHour *int32 `json:"interval_hour,omitempty"`
}

func (o ModifyComparePolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyComparePolicyReq struct{}"
	}

	return strings.Join([]string{"ModifyComparePolicyReq", string(data)}, " ")
}

type ModifyComparePolicyReqComparePolicy struct {
	value string
}

type ModifyComparePolicyReqComparePolicyEnum struct {
	NORMAL      ModifyComparePolicyReqComparePolicy
	MANY_TO_ONE ModifyComparePolicyReqComparePolicy
}

func GetModifyComparePolicyReqComparePolicyEnum() ModifyComparePolicyReqComparePolicyEnum {
	return ModifyComparePolicyReqComparePolicyEnum{
		NORMAL: ModifyComparePolicyReqComparePolicy{
			value: "normal",
		},
		MANY_TO_ONE: ModifyComparePolicyReqComparePolicy{
			value: "manyToOne",
		},
	}
}

func (c ModifyComparePolicyReqComparePolicy) Value() string {
	return c.value
}

func (c ModifyComparePolicyReqComparePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyComparePolicyReqComparePolicy) UnmarshalJSON(b []byte) error {
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
