package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageAccountInfo 用户详细信息。
type ImageAccountInfo struct {

	// 用户(组)。
	Account string `json:"account"`

	// 用户类型： * `USER` - 用户 * `USER_GROUP` - 用户组
	AccountType ImageAccountInfoAccountType `json:"account_type"`

	// 域名城。
	Domain *string `json:"domain,omitempty"`
}

func (o ImageAccountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageAccountInfo struct{}"
	}

	return strings.Join([]string{"ImageAccountInfo", string(data)}, " ")
}

type ImageAccountInfoAccountType struct {
	value string
}

type ImageAccountInfoAccountTypeEnum struct {
	USER       ImageAccountInfoAccountType
	USER_GROUP ImageAccountInfoAccountType
}

func GetImageAccountInfoAccountTypeEnum() ImageAccountInfoAccountTypeEnum {
	return ImageAccountInfoAccountTypeEnum{
		USER: ImageAccountInfoAccountType{
			value: "USER",
		},
		USER_GROUP: ImageAccountInfoAccountType{
			value: "USER_GROUP",
		},
	}
}

func (c ImageAccountInfoAccountType) Value() string {
	return c.value
}

func (c ImageAccountInfoAccountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageAccountInfoAccountType) UnmarshalJSON(b []byte) error {
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
