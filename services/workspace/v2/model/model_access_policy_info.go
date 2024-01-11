package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessPolicyInfo struct {

	// 策略名，当前只支持专线接入策略名。 * PRIVATE_ACCESS： 专线接入
	PolicyName string `json:"policy_name"`

	// 黑名单类型，当前黑名单只支持互联网。 * INTERNET： 互联网
	BlacklistType AccessPolicyInfoBlacklistType `json:"blacklist_type"`
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
