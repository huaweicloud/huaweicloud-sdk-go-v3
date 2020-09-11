/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 创建密钥对请求体
type CreateKeypairRequestBody struct {
	Keypair *CreateKeypairAction `json:"keypair"`
}

func (o CreateKeypairRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateKeypairRequestBody", string(data)}, " ")
}
