package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessPolicyDetailInfo struct {

	// 策略名。
	PolicyName *string `json:"policy_name,omitempty"`

	// 黑名单类型，当前黑名单只支持互联网。 * INTERNET： 互联网
	BlacklistType *AccessPolicyDetailInfoBlacklistType `json:"blacklist_type,omitempty"`

	// 访问控制类型。默认为接入类型 * ACCESS_TYPE： 接入类型 * IP_WHITE_LIST： IP白名单
	AccessControlType *AccessPolicyDetailInfoAccessControlType `json:"access_control_type,omitempty"`

	// 策略的ip列表。
	IpList *[]IpInfo `json:"ip_list,omitempty"`

	// IP白名单是否生效。只能单独更新，此值的优先级最高，传此值只修改该策略是否生效。
	IsEnable *bool `json:"is_enable,omitempty"`

	// IP白名单是否禁止所有Ip接入。is_enable为false时，无法修改此值。此值也只能单独更新。
	IsBlockAll *bool `json:"is_block_all,omitempty"`

	// 策略总数。
	IpTotalCount *int32 `json:"ip_total_count,omitempty"`

	// 策略id。
	PolicyId *string `json:"policy_id,omitempty"`

	// 接入策略创建的时间。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o AccessPolicyDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyDetailInfo struct{}"
	}

	return strings.Join([]string{"AccessPolicyDetailInfo", string(data)}, " ")
}

type AccessPolicyDetailInfoBlacklistType struct {
	value string
}

type AccessPolicyDetailInfoBlacklistTypeEnum struct {
	INTERNET AccessPolicyDetailInfoBlacklistType
}

func GetAccessPolicyDetailInfoBlacklistTypeEnum() AccessPolicyDetailInfoBlacklistTypeEnum {
	return AccessPolicyDetailInfoBlacklistTypeEnum{
		INTERNET: AccessPolicyDetailInfoBlacklistType{
			value: "INTERNET",
		},
	}
}

func (c AccessPolicyDetailInfoBlacklistType) Value() string {
	return c.value
}

func (c AccessPolicyDetailInfoBlacklistType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyDetailInfoBlacklistType) UnmarshalJSON(b []byte) error {
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

type AccessPolicyDetailInfoAccessControlType struct {
	value string
}

type AccessPolicyDetailInfoAccessControlTypeEnum struct {
	ACCESS_TYPE   AccessPolicyDetailInfoAccessControlType
	IP_WHITE_LIST AccessPolicyDetailInfoAccessControlType
}

func GetAccessPolicyDetailInfoAccessControlTypeEnum() AccessPolicyDetailInfoAccessControlTypeEnum {
	return AccessPolicyDetailInfoAccessControlTypeEnum{
		ACCESS_TYPE: AccessPolicyDetailInfoAccessControlType{
			value: "ACCESS_TYPE",
		},
		IP_WHITE_LIST: AccessPolicyDetailInfoAccessControlType{
			value: "IP_WHITE_LIST",
		},
	}
}

func (c AccessPolicyDetailInfoAccessControlType) Value() string {
	return c.value
}

func (c AccessPolicyDetailInfoAccessControlType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyDetailInfoAccessControlType) UnmarshalJSON(b []byte) error {
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
