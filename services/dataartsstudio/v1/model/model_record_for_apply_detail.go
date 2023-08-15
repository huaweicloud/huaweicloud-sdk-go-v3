package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RecordForApplyDetail struct {

	// 申请编号
	Id *string `json:"id,omitempty"`

	// 申请状态
	ApiApplyStatus *RecordForApplyDetailApiApplyStatus `json:"api_apply_status,omitempty"`

	// 申请类型
	ApiApplyType *RecordForApplyDetailApiApplyType `json:"api_apply_type,omitempty"`

	// api编号
	ApiId *string `json:"api_id,omitempty"`

	// api名称
	ApiName *string `json:"api_name,omitempty"`

	// 使用截止时间
	ApiUsingTime *int64 `json:"api_using_time,omitempty"`

	// app编号
	AppId *string `json:"app_id,omitempty"`

	// app名称
	AppName *string `json:"app_name,omitempty"`

	// 申请时间
	ApplyTime *int64 `json:"apply_time,omitempty"`

	// 授权时间
	ApprovalTime *int64 `json:"approval_time,omitempty"`

	// 审核人名称
	ApproverName *string `json:"approver_name,omitempty"`

	// 审核评论
	Comment *string `json:"comment,omitempty"`

	// 申请人姓名
	UserName *string `json:"user_name,omitempty"`
}

func (o RecordForApplyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordForApplyDetail struct{}"
	}

	return strings.Join([]string{"RecordForApplyDetail", string(data)}, " ")
}

type RecordForApplyDetailApiApplyStatus struct {
	value string
}

type RecordForApplyDetailApiApplyStatusEnum struct {
	STATUS_TYPE_PENDING_APPROVAL    RecordForApplyDetailApiApplyStatus
	STATUS_TYPE_REJECTED            RecordForApplyDetailApiApplyStatus
	STATUS_TYPE_PENDING_CHECK       RecordForApplyDetailApiApplyStatus
	STATUS_TYPE_PENDING_EXECUTE     RecordForApplyDetailApiApplyStatus
	STATUS_TYPE_SYNCHRONOUS_EXECUTE RecordForApplyDetailApiApplyStatus
	STATUS_TYPE_FORCED_CANCEL       RecordForApplyDetailApiApplyStatus
	STATUS_TYPE_PASSED              RecordForApplyDetailApiApplyStatus
}

func GetRecordForApplyDetailApiApplyStatusEnum() RecordForApplyDetailApiApplyStatusEnum {
	return RecordForApplyDetailApiApplyStatusEnum{
		STATUS_TYPE_PENDING_APPROVAL: RecordForApplyDetailApiApplyStatus{
			value: "STATUS_TYPE_PENDING_APPROVAL",
		},
		STATUS_TYPE_REJECTED: RecordForApplyDetailApiApplyStatus{
			value: "STATUS_TYPE_REJECTED",
		},
		STATUS_TYPE_PENDING_CHECK: RecordForApplyDetailApiApplyStatus{
			value: "STATUS_TYPE_PENDING_CHECK",
		},
		STATUS_TYPE_PENDING_EXECUTE: RecordForApplyDetailApiApplyStatus{
			value: "STATUS_TYPE_PENDING_EXECUTE",
		},
		STATUS_TYPE_SYNCHRONOUS_EXECUTE: RecordForApplyDetailApiApplyStatus{
			value: "STATUS_TYPE_SYNCHRONOUS_EXECUTE",
		},
		STATUS_TYPE_FORCED_CANCEL: RecordForApplyDetailApiApplyStatus{
			value: "STATUS_TYPE_FORCED_CANCEL",
		},
		STATUS_TYPE_PASSED: RecordForApplyDetailApiApplyStatus{
			value: "STATUS_TYPE_PASSED",
		},
	}
}

func (c RecordForApplyDetailApiApplyStatus) Value() string {
	return c.value
}

func (c RecordForApplyDetailApiApplyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordForApplyDetailApiApplyStatus) UnmarshalJSON(b []byte) error {
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

type RecordForApplyDetailApiApplyType struct {
	value string
}

type RecordForApplyDetailApiApplyTypeEnum struct {
	APPLY_TYPE_PUBLISH              RecordForApplyDetailApiApplyType
	APPLY_TYPE_AUTHORIZE            RecordForApplyDetailApiApplyType
	APPLY_TYPE_APPLY                RecordForApplyDetailApiApplyType
	APPLY_TYPE_RENEW                RecordForApplyDetailApiApplyType
	APPLY_TYPE_STOP                 RecordForApplyDetailApiApplyType
	APPLY_TYPE_RECOVER              RecordForApplyDetailApiApplyType
	APPLY_TYPE_API_CANCEL_AUTHORIZE RecordForApplyDetailApiApplyType
	APPLY_TYPE_APP_CANCEL_AUTHORIZE RecordForApplyDetailApiApplyType
	APPLY_TYPE_OFFLINE              RecordForApplyDetailApiApplyType
}

func GetRecordForApplyDetailApiApplyTypeEnum() RecordForApplyDetailApiApplyTypeEnum {
	return RecordForApplyDetailApiApplyTypeEnum{
		APPLY_TYPE_PUBLISH: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_PUBLISH",
		},
		APPLY_TYPE_AUTHORIZE: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_AUTHORIZE",
		},
		APPLY_TYPE_APPLY: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_APPLY",
		},
		APPLY_TYPE_RENEW: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_RENEW",
		},
		APPLY_TYPE_STOP: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_STOP",
		},
		APPLY_TYPE_RECOVER: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_RECOVER",
		},
		APPLY_TYPE_API_CANCEL_AUTHORIZE: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_API_CANCEL_AUTHORIZE",
		},
		APPLY_TYPE_APP_CANCEL_AUTHORIZE: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_APP_CANCEL_AUTHORIZE",
		},
		APPLY_TYPE_OFFLINE: RecordForApplyDetailApiApplyType{
			value: "APPLY_TYPE_OFFLINE",
		},
	}
}

func (c RecordForApplyDetailApiApplyType) Value() string {
	return c.value
}

func (c RecordForApplyDetailApiApplyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordForApplyDetailApiApplyType) UnmarshalJSON(b []byte) error {
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
