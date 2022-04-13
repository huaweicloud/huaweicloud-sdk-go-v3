package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ListKeysRequestBody struct {
	// 指定查询返回记录条数，如果指定查询记录条数小于存在的条数，响应参数“truncated”将返回“true”，表示存在分页。取值在密钥最大个数范围以内。例如：100

	Limit *string `json:"limit,omitempty"`
	// 分页查询起始位置标识。分页查询收到的响应参数“truncated”为“true”时，可以发送连续的请求获取更多的记录条数，“marker”设置为响应的next_marker的值。例如：10

	Marker *string `json:"marker,omitempty"`
	// 密钥状态，满足正则匹配“^[1-5]{1}$”，枚举如下：  - “1”表示待激活状态  - “2”表示启用状态  - “3”表示禁用状态  - “4”表示计划删除状态  - “5”表示等待导入状态

	KeyState *string `json:"key_state,omitempty"`
	// 密钥生成算法，默认为“AES_256”。查询所有（包含非对称）密钥需要指定参数“ALL”。  - AES_256  - SM4  - RSA_2048  - RSA_3072  - RSA_4096  - EC_P256  - EC_P384  - SM2  - ALL

	KeySpec *ListKeysRequestBodyKeySpec `json:"key_spec,omitempty"`
	// 请求消息序列号，36字节序列号。 例如：919c82d4-8046-4722-9094-35c3c6524cff

	Sequence *string `json:"sequence,omitempty"`
}

func (o ListKeysRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeysRequestBody struct{}"
	}

	return strings.Join([]string{"ListKeysRequestBody", string(data)}, " ")
}

type ListKeysRequestBodyKeySpec struct {
	value string
}

type ListKeysRequestBodyKeySpecEnum struct {
	AES_256  ListKeysRequestBodyKeySpec
	SM4      ListKeysRequestBodyKeySpec
	RSA_2048 ListKeysRequestBodyKeySpec
	RSA_3072 ListKeysRequestBodyKeySpec
	RSA_4096 ListKeysRequestBodyKeySpec
	EC_P256  ListKeysRequestBodyKeySpec
	EC_P384  ListKeysRequestBodyKeySpec
	SM2      ListKeysRequestBodyKeySpec
	ALL      ListKeysRequestBodyKeySpec
}

func GetListKeysRequestBodyKeySpecEnum() ListKeysRequestBodyKeySpecEnum {
	return ListKeysRequestBodyKeySpecEnum{
		AES_256: ListKeysRequestBodyKeySpec{
			value: "AES_256",
		},
		SM4: ListKeysRequestBodyKeySpec{
			value: "SM4",
		},
		RSA_2048: ListKeysRequestBodyKeySpec{
			value: "RSA_2048",
		},
		RSA_3072: ListKeysRequestBodyKeySpec{
			value: "RSA_3072",
		},
		RSA_4096: ListKeysRequestBodyKeySpec{
			value: "RSA_4096",
		},
		EC_P256: ListKeysRequestBodyKeySpec{
			value: "EC_P256",
		},
		EC_P384: ListKeysRequestBodyKeySpec{
			value: "EC_P384",
		},
		SM2: ListKeysRequestBodyKeySpec{
			value: "SM2",
		},
		ALL: ListKeysRequestBodyKeySpec{
			value: "ALL",
		},
	}
}

func (c ListKeysRequestBodyKeySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListKeysRequestBodyKeySpec) UnmarshalJSON(b []byte) error {
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
