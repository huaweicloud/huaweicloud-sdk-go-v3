package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigReq 接入配置
type AccessConfigReq struct {

	// 接入方式。 - INTERNET：表示Internet接入。 - DEDICATED：表示专线接入。 - BOTH：表示两种接入方式都支持。
	AccessMode AccessConfigReqAccessMode `json:"access_mode"`

	// 专线接入网段列表，多个网段信息用分号隔开，列表长度不超过5。
	DedicatedCidrs *string `json:"dedicated_cidrs,omitempty"`
}

func (o AccessConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigReq struct{}"
	}

	return strings.Join([]string{"AccessConfigReq", string(data)}, " ")
}

type AccessConfigReqAccessMode struct {
	value string
}

type AccessConfigReqAccessModeEnum struct {
	INTERNET  AccessConfigReqAccessMode
	DEDICATED AccessConfigReqAccessMode
	BOTH      AccessConfigReqAccessMode
}

func GetAccessConfigReqAccessModeEnum() AccessConfigReqAccessModeEnum {
	return AccessConfigReqAccessModeEnum{
		INTERNET: AccessConfigReqAccessMode{
			value: "INTERNET",
		},
		DEDICATED: AccessConfigReqAccessMode{
			value: "DEDICATED",
		},
		BOTH: AccessConfigReqAccessMode{
			value: "BOTH",
		},
	}
}

func (c AccessConfigReqAccessMode) Value() string {
	return c.value
}

func (c AccessConfigReqAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigReqAccessMode) UnmarshalJSON(b []byte) error {
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
