package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// Request Object
type ListScalingActivityV2LogsRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Type *ListScalingActivityV2LogsRequestType `json:"type,omitempty"`

	Status *ListScalingActivityV2LogsRequestStatus `json:"status,omitempty"`
}

func (o ListScalingActivityV2LogsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListScalingActivityV2LogsRequest struct{}"
	}

	return strings.Join([]string{"ListScalingActivityV2LogsRequest", string(data)}, " ")
}

type ListScalingActivityV2LogsRequestType struct {
	value string
}

type ListScalingActivityV2LogsRequestTypeEnum struct {
	NORMAL             ListScalingActivityV2LogsRequestType
	MANNUAL_REMOVE     ListScalingActivityV2LogsRequestType
	MANNUAL_DELETE     ListScalingActivityV2LogsRequestType
	MANNUAL_ADD        ListScalingActivityV2LogsRequestType
	ELB_CHECK_DELETE   ListScalingActivityV2LogsRequestType
	AUDIT_CHECK_DELETE ListScalingActivityV2LogsRequestType
	MODIFY_ELB         ListScalingActivityV2LogsRequestType
}

func GetListScalingActivityV2LogsRequestTypeEnum() ListScalingActivityV2LogsRequestTypeEnum {
	return ListScalingActivityV2LogsRequestTypeEnum{
		NORMAL: ListScalingActivityV2LogsRequestType{
			value: "NORMAL",
		},
		MANNUAL_REMOVE: ListScalingActivityV2LogsRequestType{
			value: "MANNUAL_REMOVE",
		},
		MANNUAL_DELETE: ListScalingActivityV2LogsRequestType{
			value: "MANNUAL_DELETE",
		},
		MANNUAL_ADD: ListScalingActivityV2LogsRequestType{
			value: "MANNUAL_ADD",
		},
		ELB_CHECK_DELETE: ListScalingActivityV2LogsRequestType{
			value: "ELB_CHECK_DELETE",
		},
		AUDIT_CHECK_DELETE: ListScalingActivityV2LogsRequestType{
			value: "AUDIT_CHECK_DELETE",
		},
		MODIFY_ELB: ListScalingActivityV2LogsRequestType{
			value: "MODIFY_ELB",
		},
	}
}

func (c ListScalingActivityV2LogsRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListScalingActivityV2LogsRequestType) UnmarshalJSON(b []byte) error {
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

type ListScalingActivityV2LogsRequestStatus struct {
	value string
}

type ListScalingActivityV2LogsRequestStatusEnum struct {
	SUCCESS ListScalingActivityV2LogsRequestStatus
	FAIL    ListScalingActivityV2LogsRequestStatus
	DOING   ListScalingActivityV2LogsRequestStatus
}

func GetListScalingActivityV2LogsRequestStatusEnum() ListScalingActivityV2LogsRequestStatusEnum {
	return ListScalingActivityV2LogsRequestStatusEnum{
		SUCCESS: ListScalingActivityV2LogsRequestStatus{
			value: "SUCCESS",
		},
		FAIL: ListScalingActivityV2LogsRequestStatus{
			value: "FAIL",
		},
		DOING: ListScalingActivityV2LogsRequestStatus{
			value: "DOING",
		},
	}
}

func (c ListScalingActivityV2LogsRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListScalingActivityV2LogsRequestStatus) UnmarshalJSON(b []byte) error {
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
