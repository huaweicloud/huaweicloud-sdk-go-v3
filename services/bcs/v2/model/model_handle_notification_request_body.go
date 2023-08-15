package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HandleNotificationRequestBody 处理联盟邀请body数据
type HandleNotificationRequestBody struct {

	// 邀请目标通道
	ChannelName string `json:"channel_name"`

	// 处理邀请
	Status HandleNotificationRequestBodyStatus `json:"status"`

	InvitorInfo *HandleNotificationInvitor `json:"invitor_info"`

	InviteeInfo *HandleNotificationInvitee `json:"invitee_info,omitempty"`

	// 加入联盟的组织，同意加入时必填
	InvitedOrgs *[]HandleNotificationOrg `json:"invited_orgs,omitempty"`
}

func (o HandleNotificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleNotificationRequestBody struct{}"
	}

	return strings.Join([]string{"HandleNotificationRequestBody", string(data)}, " ")
}

type HandleNotificationRequestBodyStatus struct {
	value string
}

type HandleNotificationRequestBodyStatusEnum struct {
	AGREED HandleNotificationRequestBodyStatus
	REJECT HandleNotificationRequestBodyStatus
}

func GetHandleNotificationRequestBodyStatusEnum() HandleNotificationRequestBodyStatusEnum {
	return HandleNotificationRequestBodyStatusEnum{
		AGREED: HandleNotificationRequestBodyStatus{
			value: "agreed",
		},
		REJECT: HandleNotificationRequestBodyStatus{
			value: "reject",
		},
	}
}

func (c HandleNotificationRequestBodyStatus) Value() string {
	return c.value
}

func (c HandleNotificationRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HandleNotificationRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
