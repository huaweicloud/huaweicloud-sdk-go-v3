package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessPolicyDetailInfo struct {

	// 策略名，当前只支持专线接入策略名。 * PRIVATE_ACCESS： 专线接入
	PolicyName AccessPolicyDetailInfoPolicyName `json:"policy_name"`

	// 黑名单类型，当前黑名单只支持互联网。 * INTERNET： 互联网
	BlacklistType AccessPolicyDetailInfoBlacklistType `json:"blacklist_type"`

	// 策略id。
	PolicyId *string `json:"policy_id,omitempty"`

	// 用户otp设备绑定时间。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o AccessPolicyDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyDetailInfo struct{}"
	}

	return strings.Join([]string{"AccessPolicyDetailInfo", string(data)}, " ")
}

type AccessPolicyDetailInfoPolicyName struct {
	value string
}

type AccessPolicyDetailInfoPolicyNameEnum struct {
	PRIVATE_ACCESS AccessPolicyDetailInfoPolicyName
}

func GetAccessPolicyDetailInfoPolicyNameEnum() AccessPolicyDetailInfoPolicyNameEnum {
	return AccessPolicyDetailInfoPolicyNameEnum{
		PRIVATE_ACCESS: AccessPolicyDetailInfoPolicyName{
			value: "PRIVATE_ACCESS",
		},
	}
}

func (c AccessPolicyDetailInfoPolicyName) Value() string {
	return c.value
}

func (c AccessPolicyDetailInfoPolicyName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyDetailInfoPolicyName) UnmarshalJSON(b []byte) error {
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
