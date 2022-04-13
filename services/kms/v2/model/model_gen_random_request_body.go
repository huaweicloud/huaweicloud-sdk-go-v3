package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenRandomRequestBody struct {
	// 随机数的bit位长度。 取值为8的倍数，取值范围为8~8192。 随机数的bit位长度，取值为“512”。

	RandomDataLength string `json:"random_data_length"`
	// 请求消息序列号，36字节序列号。 例如：919c82d4-8046-4722-9094-35c3c6524cff

	Sequence *string `json:"sequence,omitempty"`
}

func (o GenRandomRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenRandomRequestBody struct{}"
	}

	return strings.Join([]string{"GenRandomRequestBody", string(data)}, " ")
}
