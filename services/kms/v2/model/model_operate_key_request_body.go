package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateKeyRequestBody struct {
	// 密钥ID，36字节，满足正则匹配“^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$”。 例如：0d0466b0-e727-4d9c-b35d-f84bb474a37f。

	KeyId string `json:"key_id"`
	// 请求消息序列号，36字节序列号。 例如：919c82d4-8046-4722-9094-35c3c6524cff

	Sequence *string `json:"sequence,omitempty"`
}

func (o OperateKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateKeyRequestBody struct{}"
	}

	return strings.Join([]string{"OperateKeyRequestBody", string(data)}, " ")
}
