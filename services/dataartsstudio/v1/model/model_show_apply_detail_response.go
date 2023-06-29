package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApplyDetailResponse Response Object
type ShowApplyDetailResponse struct {

	// 申请编号
	Id *string `json:"id,omitempty"`

	// 申请状态
	ApiApplyStatus *ShowApplyDetailResponseApiApplyStatus `json:"api_apply_status,omitempty"`

	// 申请类型
	ApiApplyType *ShowApplyDetailResponseApiApplyType `json:"api_apply_type,omitempty"`

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
	UserName       *string `json:"user_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowApplyDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplyDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowApplyDetailResponse", string(data)}, " ")
}

type ShowApplyDetailResponseApiApplyStatus struct {
	value string
}

type ShowApplyDetailResponseApiApplyStatusEnum struct {
	STATUS_TYPE_PENDING_APPROVAL    ShowApplyDetailResponseApiApplyStatus
	STATUS_TYPE_REJECTED            ShowApplyDetailResponseApiApplyStatus
	STATUS_TYPE_PENDING_CHECK       ShowApplyDetailResponseApiApplyStatus
	STATUS_TYPE_PENDING_EXECUTE     ShowApplyDetailResponseApiApplyStatus
	STATUS_TYPE_SYNCHRONOUS_EXECUTE ShowApplyDetailResponseApiApplyStatus
	STATUS_TYPE_FORCED_CANCEL       ShowApplyDetailResponseApiApplyStatus
	STATUS_TYPE_PASSED              ShowApplyDetailResponseApiApplyStatus
}

func GetShowApplyDetailResponseApiApplyStatusEnum() ShowApplyDetailResponseApiApplyStatusEnum {
	return ShowApplyDetailResponseApiApplyStatusEnum{
		STATUS_TYPE_PENDING_APPROVAL: ShowApplyDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_PENDING_APPROVAL",
		},
		STATUS_TYPE_REJECTED: ShowApplyDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_REJECTED",
		},
		STATUS_TYPE_PENDING_CHECK: ShowApplyDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_PENDING_CHECK",
		},
		STATUS_TYPE_PENDING_EXECUTE: ShowApplyDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_PENDING_EXECUTE",
		},
		STATUS_TYPE_SYNCHRONOUS_EXECUTE: ShowApplyDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_SYNCHRONOUS_EXECUTE",
		},
		STATUS_TYPE_FORCED_CANCEL: ShowApplyDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_FORCED_CANCEL",
		},
		STATUS_TYPE_PASSED: ShowApplyDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_PASSED",
		},
	}
}

func (c ShowApplyDetailResponseApiApplyStatus) Value() string {
	return c.value
}

func (c ShowApplyDetailResponseApiApplyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApplyDetailResponseApiApplyStatus) UnmarshalJSON(b []byte) error {
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

type ShowApplyDetailResponseApiApplyType struct {
	value string
}

type ShowApplyDetailResponseApiApplyTypeEnum struct {
	APPLY_TYPE_PUBLISH              ShowApplyDetailResponseApiApplyType
	APPLY_TYPE_AUTHORIZE            ShowApplyDetailResponseApiApplyType
	APPLY_TYPE_APPLY                ShowApplyDetailResponseApiApplyType
	APPLY_TYPE_RENEW                ShowApplyDetailResponseApiApplyType
	APPLY_TYPE_STOP                 ShowApplyDetailResponseApiApplyType
	APPLY_TYPE_RECOVER              ShowApplyDetailResponseApiApplyType
	APPLY_TYPE_API_CANCEL_AUTHORIZE ShowApplyDetailResponseApiApplyType
	APPLY_TYPE_APP_CANCEL_AUTHORIZE ShowApplyDetailResponseApiApplyType
	APPLY_TYPE_OFFLINE              ShowApplyDetailResponseApiApplyType
}

func GetShowApplyDetailResponseApiApplyTypeEnum() ShowApplyDetailResponseApiApplyTypeEnum {
	return ShowApplyDetailResponseApiApplyTypeEnum{
		APPLY_TYPE_PUBLISH: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_PUBLISH",
		},
		APPLY_TYPE_AUTHORIZE: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_AUTHORIZE",
		},
		APPLY_TYPE_APPLY: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_APPLY",
		},
		APPLY_TYPE_RENEW: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_RENEW",
		},
		APPLY_TYPE_STOP: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_STOP",
		},
		APPLY_TYPE_RECOVER: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_RECOVER",
		},
		APPLY_TYPE_API_CANCEL_AUTHORIZE: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_API_CANCEL_AUTHORIZE",
		},
		APPLY_TYPE_APP_CANCEL_AUTHORIZE: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_APP_CANCEL_AUTHORIZE",
		},
		APPLY_TYPE_OFFLINE: ShowApplyDetailResponseApiApplyType{
			value: "APPLY_TYPE_OFFLINE",
		},
	}
}

func (c ShowApplyDetailResponseApiApplyType) Value() string {
	return c.value
}

func (c ShowApplyDetailResponseApiApplyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApplyDetailResponseApiApplyType) UnmarshalJSON(b []byte) error {
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
