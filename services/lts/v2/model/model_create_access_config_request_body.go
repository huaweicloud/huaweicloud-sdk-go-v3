package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAccessConfigRequestBody 创建日志接入请求体
type CreateAccessConfigRequestBody struct {

	// 日志接入名称。 满足正则表达式：^(?!\\.)(?!_)(?!.*?\\.$)[\\u4e00-\\u9fa5a-zA-Z0-9-_.]{1,64}$
	AccessConfigName string `json:"access_config_name"`

	// 日志接入类型。AGENT：ECS接入,K8S_CCE:CCE接入
	AccessConfigType CreateAccessConfigRequestBodyAccessConfigType `json:"access_config_type"`

	AccessConfigDetail *AccessConfigDeatilCreate `json:"access_config_detail"`

	LogInfo *AccessConfigBaseLogInfoCreate `json:"log_info"`

	HostGroupInfo *AccessConfigHostGroupIdListCreate `json:"host_group_info,omitempty"`

	// 标签信息。KEY不能重复,最多20个标签
	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`

	// 二进制采集
	BinaryCollect *bool `json:"binary_collect,omitempty"`

	// 日志拆分
	LogSplit *bool `json:"log_split,omitempty"`
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
	AGENT    CreateAccessConfigRequestBodyAccessConfigType
	K8_S_CCE CreateAccessConfigRequestBodyAccessConfigType
}

func GetCreateAccessConfigRequestBodyAccessConfigTypeEnum() CreateAccessConfigRequestBodyAccessConfigTypeEnum {
	return CreateAccessConfigRequestBodyAccessConfigTypeEnum{
		AGENT: CreateAccessConfigRequestBodyAccessConfigType{
			value: "AGENT",
		},
		K8_S_CCE: CreateAccessConfigRequestBodyAccessConfigType{
			value: "K8S_CCE",
		},
	}
}

func (c CreateAccessConfigRequestBodyAccessConfigType) Value() string {
	return c.value
}

func (c CreateAccessConfigRequestBodyAccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccessConfigRequestBodyAccessConfigType) UnmarshalJSON(b []byte) error {
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
