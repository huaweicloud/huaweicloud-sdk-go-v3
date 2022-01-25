package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建日志接入请求体
type CreateAccessConfigRequestBody struct {
	// 日志接入名称

	AccessConfigName string `json:"access_config_name"`
	// 日志接入类型。AGENT：主机接入类型

	AccessConfigType CreateAccessConfigRequestBodyAccessConfigType `json:"access_config_type"`

	AccessConfigDetail *AccessConfigDeatilCreate `json:"access_config_detail"`

	LogInfo *AccessConfigBaseLogInfoCreate `json:"log_info"`

	HostGroupInfo *AccessConfigHostGroupIdListCreate `json:"host_group_info,omitempty"`

	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`
}

func (o CreateAccessConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAccessConfigRequestBody", string(data)}, " ")
}

type CreateAccessConfigRequestBodyAccessConfigType struct {
	value string
}

type CreateAccessConfigRequestBodyAccessConfigTypeEnum struct {
	AGENT CreateAccessConfigRequestBodyAccessConfigType
}

func GetCreateAccessConfigRequestBodyAccessConfigTypeEnum() CreateAccessConfigRequestBodyAccessConfigTypeEnum {
	return CreateAccessConfigRequestBodyAccessConfigTypeEnum{
		AGENT: CreateAccessConfigRequestBodyAccessConfigType{
			value: "AGENT",
		},
	}
}

func (c CreateAccessConfigRequestBodyAccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccessConfigRequestBodyAccessConfigType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
