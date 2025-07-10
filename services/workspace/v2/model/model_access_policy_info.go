package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessPolicyInfo struct {

	// 策略名。
	PolicyName *string `json:"policy_name,omitempty"`

	// 黑名单类型，当前黑名单只支持互联网。 * INTERNET： 互联网
	BlacklistType *AccessPolicyInfoBlacklistType `json:"blacklist_type,omitempty"`

	// 访问控制类型。默认为接入类型 * ACCESS_TYPE： 接入类型 * IP_WHITE_LIST： IP白名单
	AccessControlType *AccessPolicyInfoAccessControlType `json:"access_control_type,omitempty"`

	// 策略的ip列表。
	IpList *[]IpInfo `json:"ip_list,omitempty"`

	// IP白名单是否生效。只能单独更新，此值的优先级最高，传此值只修改该策略是否生效。
	IsEnable *bool `json:"is_enable,omitempty"`

	// IP白名单是否禁止所有Ip接入。is_enable为false时，无法修改此值。此值也只能单独更新。
	IsBlockAll *bool `json:"is_block_all,omitempty"`

	// 策略总数。
	IpTotalCount *int32 `json:"ip_total_count,omitempty"`
}

func (o AccessPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyInfo struct{}"
	}

	return strings.Join([]string{"AccessPolicyInfo", string(data)}, " ")
}

type AccessPolicyInfoBlacklistType struct {
	value string
}

type AccessPolicyInfoBlacklistTypeEnum struct {
	INTERNET AccessPolicyInfoBlacklistType
}

func GetAccessPolicyInfoBlacklistTypeEnum() AccessPolicyInfoBlacklistTypeEnum {
	return AccessPolicyInfoBlacklistTypeEnum{
		INTERNET: AccessPolicyInfoBlacklistType{
			value: "INTERNET",
		},
	}
}

func (c AccessPolicyInfoBlacklistType) Value() string {
	return c.value
}

func (c AccessPolicyInfoBlacklistType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyInfoBlacklistType) UnmarshalJSON(b []byte) error {
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

type AccessPolicyInfoAccessControlType struct {
	value string
}

type AccessPolicyInfoAccessControlTypeEnum struct {
	ACCESS_TYPE   AccessPolicyInfoAccessControlType
	IP_WHITE_LIST AccessPolicyInfoAccessControlType
}

func GetAccessPolicyInfoAccessControlTypeEnum() AccessPolicyInfoAccessControlTypeEnum {
	return AccessPolicyInfoAccessControlTypeEnum{
		ACCESS_TYPE: AccessPolicyInfoAccessControlType{
			value: "ACCESS_TYPE",
		},
		IP_WHITE_LIST: AccessPolicyInfoAccessControlType{
			value: "IP_WHITE_LIST",
		},
	}
}

func (c AccessPolicyInfoAccessControlType) Value() string {
	return c.value
}

func (c AccessPolicyInfoAccessControlType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyInfoAccessControlType) UnmarshalJSON(b []byte) error {
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
