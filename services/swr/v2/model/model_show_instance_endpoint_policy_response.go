package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceEndpointPolicyResponse Response Object
type ShowInstanceEndpointPolicyResponse struct {

	// 公网访问状态，Enable、Enabling、EnableFailed、Disable、Disabling、DisableFailed
	Status *ShowInstanceEndpointPolicyResponseStatus `json:"status,omitempty"`

	// 白名单列表
	IpList         *[]IpInfo `json:"ip_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowInstanceEndpointPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceEndpointPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceEndpointPolicyResponse", string(data)}, " ")
}

type ShowInstanceEndpointPolicyResponseStatus struct {
	value string
}

type ShowInstanceEndpointPolicyResponseStatusEnum struct {
	ENABLING       ShowInstanceEndpointPolicyResponseStatus
	ENABLE         ShowInstanceEndpointPolicyResponseStatus
	ENABLE_FAILED  ShowInstanceEndpointPolicyResponseStatus
	DISABLE        ShowInstanceEndpointPolicyResponseStatus
	DISABLING      ShowInstanceEndpointPolicyResponseStatus
	DISABLE_FAILED ShowInstanceEndpointPolicyResponseStatus
}

func GetShowInstanceEndpointPolicyResponseStatusEnum() ShowInstanceEndpointPolicyResponseStatusEnum {
	return ShowInstanceEndpointPolicyResponseStatusEnum{
		ENABLING: ShowInstanceEndpointPolicyResponseStatus{
			value: "Enabling",
		},
		ENABLE: ShowInstanceEndpointPolicyResponseStatus{
			value: "Enable",
		},
		ENABLE_FAILED: ShowInstanceEndpointPolicyResponseStatus{
			value: "EnableFailed",
		},
		DISABLE: ShowInstanceEndpointPolicyResponseStatus{
			value: "Disable",
		},
		DISABLING: ShowInstanceEndpointPolicyResponseStatus{
			value: "Disabling",
		},
		DISABLE_FAILED: ShowInstanceEndpointPolicyResponseStatus{
			value: "DisableFailed",
		},
	}
}

func (c ShowInstanceEndpointPolicyResponseStatus) Value() string {
	return c.value
}

func (c ShowInstanceEndpointPolicyResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceEndpointPolicyResponseStatus) UnmarshalJSON(b []byte) error {
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
