package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMessageDetailResponse Response Object
type ShowMessageDetailResponse struct {

	// 申请编号
	Id *string `json:"id,omitempty"`

	// 申请状态
	ApiApplyStatus *ShowMessageDetailResponseApiApplyStatus `json:"api_apply_status,omitempty"`

	// 申请类型
	ApiApplyType *ShowMessageDetailResponseApiApplyType `json:"api_apply_type,omitempty"`

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

func (o ShowMessageDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMessageDetailResponse", string(data)}, " ")
}

type ShowMessageDetailResponseApiApplyStatus struct {
	value string
}

type ShowMessageDetailResponseApiApplyStatusEnum struct {
	STATUS_TYPE_PENDING_APPROVAL    ShowMessageDetailResponseApiApplyStatus
	STATUS_TYPE_REJECTED            ShowMessageDetailResponseApiApplyStatus
	STATUS_TYPE_PENDING_CHECK       ShowMessageDetailResponseApiApplyStatus
	STATUS_TYPE_PENDING_EXECUTE     ShowMessageDetailResponseApiApplyStatus
	STATUS_TYPE_SYNCHRONOUS_EXECUTE ShowMessageDetailResponseApiApplyStatus
	STATUS_TYPE_FORCED_CANCEL       ShowMessageDetailResponseApiApplyStatus
	STATUS_TYPE_PASSED              ShowMessageDetailResponseApiApplyStatus
}

func GetShowMessageDetailResponseApiApplyStatusEnum() ShowMessageDetailResponseApiApplyStatusEnum {
	return ShowMessageDetailResponseApiApplyStatusEnum{
		STATUS_TYPE_PENDING_APPROVAL: ShowMessageDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_PENDING_APPROVAL",
		},
		STATUS_TYPE_REJECTED: ShowMessageDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_REJECTED",
		},
		STATUS_TYPE_PENDING_CHECK: ShowMessageDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_PENDING_CHECK",
		},
		STATUS_TYPE_PENDING_EXECUTE: ShowMessageDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_PENDING_EXECUTE",
		},
		STATUS_TYPE_SYNCHRONOUS_EXECUTE: ShowMessageDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_SYNCHRONOUS_EXECUTE",
		},
		STATUS_TYPE_FORCED_CANCEL: ShowMessageDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_FORCED_CANCEL",
		},
		STATUS_TYPE_PASSED: ShowMessageDetailResponseApiApplyStatus{
			value: "STATUS_TYPE_PASSED",
		},
	}
}

func (c ShowMessageDetailResponseApiApplyStatus) Value() string {
	return c.value
}

func (c ShowMessageDetailResponseApiApplyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMessageDetailResponseApiApplyStatus) UnmarshalJSON(b []byte) error {
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

type ShowMessageDetailResponseApiApplyType struct {
	value string
}

type ShowMessageDetailResponseApiApplyTypeEnum struct {
	APPLY_TYPE_PUBLISH              ShowMessageDetailResponseApiApplyType
	APPLY_TYPE_AUTHORIZE            ShowMessageDetailResponseApiApplyType
	APPLY_TYPE_APPLY                ShowMessageDetailResponseApiApplyType
	APPLY_TYPE_RENEW                ShowMessageDetailResponseApiApplyType
	APPLY_TYPE_STOP                 ShowMessageDetailResponseApiApplyType
	APPLY_TYPE_RECOVER              ShowMessageDetailResponseApiApplyType
	APPLY_TYPE_API_CANCEL_AUTHORIZE ShowMessageDetailResponseApiApplyType
	APPLY_TYPE_APP_CANCEL_AUTHORIZE ShowMessageDetailResponseApiApplyType
	APPLY_TYPE_OFFLINE              ShowMessageDetailResponseApiApplyType
}

func GetShowMessageDetailResponseApiApplyTypeEnum() ShowMessageDetailResponseApiApplyTypeEnum {
	return ShowMessageDetailResponseApiApplyTypeEnum{
		APPLY_TYPE_PUBLISH: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_PUBLISH",
		},
		APPLY_TYPE_AUTHORIZE: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_AUTHORIZE",
		},
		APPLY_TYPE_APPLY: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_APPLY",
		},
		APPLY_TYPE_RENEW: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_RENEW",
		},
		APPLY_TYPE_STOP: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_STOP",
		},
		APPLY_TYPE_RECOVER: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_RECOVER",
		},
		APPLY_TYPE_API_CANCEL_AUTHORIZE: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_API_CANCEL_AUTHORIZE",
		},
		APPLY_TYPE_APP_CANCEL_AUTHORIZE: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_APP_CANCEL_AUTHORIZE",
		},
		APPLY_TYPE_OFFLINE: ShowMessageDetailResponseApiApplyType{
			value: "APPLY_TYPE_OFFLINE",
		},
	}
}

func (c ShowMessageDetailResponseApiApplyType) Value() string {
	return c.value
}

func (c ShowMessageDetailResponseApiApplyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMessageDetailResponseApiApplyType) UnmarshalJSON(b []byte) error {
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
