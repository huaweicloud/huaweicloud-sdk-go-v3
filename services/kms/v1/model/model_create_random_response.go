/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateRandomResponse struct {
	// 随机数16进制表示，两位表示1byte。随机数的长度与用户传入的参数 “random_data_length”的长度保持一致。
	RandomData *string `json:"random_data,omitempty"`
}

func (o CreateRandomResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateRandomResponse", string(data)}, " ")
}
